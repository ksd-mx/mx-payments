variable "env" {
    description = "Environment name."
    type = string
}

variable "eks_name" {
    description = "Name of the cluster."
    type = string
}

variable "enable_cluster_autoscaler" {
    description = "Enables the deployment of the cluster autoscaler."
    type = bool
    default= false
}

variable "cluster_autoscaler_helm_version" {
    description = "Cluster autoscaler Helm version."
    type = string
}

variable "openid_provider_arn" {
    description = "IAM OpenId Connect Provider ARN."
    type = string
}