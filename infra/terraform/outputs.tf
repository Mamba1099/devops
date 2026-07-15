output "cluster_endpoint" {
  description = "EKS cluster endpoint"
  value       = aws_eks_cluster.main.endpoint
}

output "cluster_name" {
  description = "EKS cluster name"
  value       = aws_eks_cluster.main.name
}

output "database_endpoint" {
  description = "RDS database endpoint"
  value       = aws_db_instance.myapp.address
}

output "s3_bucket_name" {
  description = "S3 bucket name"
  value       = aws_s3_bucket.myapp.bucket
}