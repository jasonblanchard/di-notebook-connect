locals {
  exec_role_arn  = "arn:aws:iam::076797644834:role/di-apilambda-task-exec-role-2862104"
  api_gateway_id = "lijpimk8ec"
  authorizer_id  = "2bqgut"
  sg_id          = "sg-0e1d7185c09def069"
  subnet_ids     = ["subnet-09520cb51ce759869", "subnet-0e85b7e5bfb5fadc4"]
}

# resource "aws_s3_bucket" "source" {
#   bucket = "di-notebook-connect-lambda-source"
# }

resource "aws_lambda_function" "api" {
  function_name    = "di-notebook-connect"
  filename         = "../../bin/lambda.zip"
  handler          = "lambda"
  role             = local.exec_role_arn
  runtime          = "go1.x"
  publish          = true
  source_code_hash = filebase64sha256("../../bin/lambda.zip")

  environment {
    variables = {}
  }

  vpc_config {
    subnet_ids         = local.subnet_ids
    security_group_ids = [local.sg_id]
  }

  depends_on = [
    aws_cloudwatch_log_group.lambda,
  ]

  lifecycle {
    ignore_changes = [
      environment
    ]
  }
}

resource "aws_cloudwatch_log_group" "lambda" {
  name              = "/aws/lambda/di-notebook-connect"
  retention_in_days = 30
}

resource "aws_lambda_alias" "api" {
  name             = "release"
  function_name    = aws_lambda_function.api.arn
  function_version = aws_lambda_function.api.version
}

resource "aws_lambda_permission" "gateway" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.api.function_name
  #   qualifier     = aws_lambda_alias.api.name
  principal  = "apigateway.amazonaws.com"
  source_arn = "arn:aws:execute-api:us-east-1:076797644834:${local.api_gateway_id}/*/*/*"
}

resource "aws_apigatewayv2_integration" "api" {
  api_id                 = local.api_gateway_id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.api.invoke_arn
  integration_method     = "POST"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "api" {
  api_id             = local.api_gateway_id
  route_key          = "ANY /notebook.v1.NotebookService/{proxy+}"
  target             = "integrations/${aws_apigatewayv2_integration.api.id}"
  authorizer_id      = local.authorizer_id
  authorization_type = "JWT"
}