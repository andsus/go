GCP Security Requirements
- USe Cloud Security Scanner, to scan application vulnerability.
- No vulnerability associated with https://owasp.org/www-project-top-ten/ and https://owasp.org/www-project-api-security/
- Pass Blueshield SQA test.
- Use Identity Aware Proxy
- If storing sensitive data use Cloud KMS for encryption/decryption
- Use IAM with custom role with TBA policy.
- Use service account for application
- Enabling monitoring and loggina API are anabled for risk and decon.
- Create deny client policy based on IPs (If needed)

## Secure Terraform State
- Removing unneessary secrets from Terraform state
- Enryption at rest
- Apply least access level
- Use tf linter to lint the tf files
- Consider using dynamic secret
  
## SAST
- Vendor provided static code analysis, or option BSCA run the static code analysis for source code
- Code review with no medium, high, critical severity

## IAM
- OAuth / SAML BSCA design pattern. 