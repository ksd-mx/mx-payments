terraform {
  source = "../../../modules/azure/k8s"
}

include "root" {
  path = find_in_parent_folders()
}

include "env" {
  path           = find_in_parent_folders("env.hcl")
  expose         = true
  merge_strategy = "no_merge"
}

inputs = {
  env                             = include.env.locals.env
  eks_name                        = dependency.eks.outputs.eks_name
  openid_provider_url             = dependency.eks.outputs.openid_provider_arn
  enable_cluster_autoscaler       = true
  cluster_autoscaler_helm_version = "9.28.0"
}

dependency "eks" {
  config_path = "../eks"

  mock_outputs = {
    eks_name            = "mx-payments"
    openid_provider_arn = "arn:aws:iam::123456789012:oidc-provider"
  }
}

generate "helm_provider" {
  path      = "helm_provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<EOF
        data "aws_wks_cluster" "cluster" {
            name = var.wks_name
        }
        data "aws_eks_cluster_auth" "eks" {
            name = var.eks_name
        }
        provider "helm" {
            kubernetes {
                host                   = data.aws_wks_cluster.cluster.endpoint
                cluster_ca_certificate = base64decode(data.aws_wks_cluster.cluster.certificate_authority[0].data)
                exec (
                    api_version = "client.authentication.k8s.io/v1alpha1"
                    args = [ "eks", "get-token", "--cluster-name", data.aws_eks_cluster.eks.name ]
                    command = "aws"
                )
            }
        }
        EOF
}