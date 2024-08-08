# GitHub Actions Workflow for Go Project

## Overview
This guide outlines the steps to set up a GitHub Actions workflow for a Go project. The workflow includes building the project, setting environment variables, and committing changes to the `README.md` file.

## Steps

1. **Set Up Workflow File**
   - Create a workflow file `.github/workflows/go.yml` with the following content:

    ```yaml
    name: Go

    on:
      schedule:
        - cron: '30 2 * * *'
      push:
        branches: [ "main" ]

    jobs:
      build:
        runs-on: ubuntu-latest
        environment: production
        steps:
          - uses: actions/checkout@v4

          - name: Set up Go
            uses: actions/setup-go@v4
            with:
              go-version: '1.22.4'

          - name: Run
            env:
              API_BASE_URL: ${{ vars.API_BASE_URL }}
            run: go run ./cmd/quote

          - name: Configure git
            run: |
              git config --global user.name "Dien Vo"
              git config --global user.email "dien.vo@andpad.co.jp"

          - name: Add README.md
            run: git add README.md

          - name: Commit changes
            run: git commit -m "Update README.md with a new quote"

          - name: Push changes
            env:
              GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
            run: git push origin main
    ```

2. **Environment Configuration**
   - Set the environment variable `API_BASE_URL` in the GitHub repository settings under `Settings -> Environments`.

3. **Go Code**
   - Ensure the Go code in `cmd/quote/main.go` fetches a quote from the API and updates `README.md`.

4. **GitHub Settings**
   - Configure GitHub Actions permissions under `Settings -> Actions -> General -> Workflow permissions`.

## Summary
This setup will automatically run the workflow daily at 09:30 UTC+7 and on every push to the `main` branch. It fetches a quote from the API, updates `README.md`, and commits the changes to the repository.