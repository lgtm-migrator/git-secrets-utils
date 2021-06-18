# Git Secrets Utils

> a simpel tool convert git secrets scan result to excel sheet

```bash
git-secrets --scan -r . 2>error_logs
git-secrets-utils convert -f error_logs -o error_logs_formatted
```