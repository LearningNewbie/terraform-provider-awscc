resource "awscc_eks_nodegroup" "eks_nodegroup" {
    cluster_name = awscc_eks_cluster.main.name
    node_role    = awscc_iam_role.eks_nodegroup_role.arn
    nodegroup_name = "eks_nodegroup"
    scaling_config {
        desired_size = 2
        max_size     = 2
        min_size     = 1
    }
    subnet_ids = [
        awscc_ec2_subnet.subnet1.id,
        awscc_ec2_subnet.subnet2.id
    ]
    tags = [
        {
          key = "Name"
          value = "eks_nodegroup"
        }

}

resource "awscc_iam_role" "eks_nodegroup_role" {
    role_name = "eks_nodegroup_role"
    assume_role_policy_document = jsonencode({
        Statement = [{
            Action = "sts:AssumeRole"
            Effect = "Allow"
            Principal = {
                Service = "ec2.amazonaws.com"
            }
        }]
        Version = "2012-10-17"
    })
    managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
      ]
}


