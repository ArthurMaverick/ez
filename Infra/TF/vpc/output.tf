output "vpc_arn" {
    value = module.vpc.vpc_arn
}

output "vpc_cdir" {
    value = module.vpc.vpc_cidr_block
}

output "private_subnets_arns" {
    value = module.vpc.private_subnet_arns
}

output "private_subnets_ids" {
    value = module.vpc.private_subnets
}

output "public_subnets_arns" {
    value = module.vpc.public_subnet_arns
}

output "public_subnets_ids" {
    value = module.vpc.public_subnets
}
