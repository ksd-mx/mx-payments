variable "env" {
  description = "Environment name."
  type        = string
}

variable "eks_version" {
  description = "Desired Kubernetes master version."
  type        = string
}

variable "eks_name" {
  description = "Name of the EKS cluster."
  type        = string
}

variable "subnet_ids" {
  description = "List of subnet IDs across a minimum of 2 azs."
  type        = list(string)
}

variable "node_iam_policies" {
  description = "List of IAM policies to attach to the node role."
  type        = map(any)
  default = {
    1 = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
    2 = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
    3 = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
    4 = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
  }
}

variable "node_groups" {
  description = "EKS node groups."
  type        = map(any)
}

variable "enable_irsa" {
  description = "Determines whether to create an OpenID Connect provider for the cluster to enable IRSA."
  type        = bool
  default     = false
}