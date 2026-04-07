---
name: "pr-creator"
description: "Helps create GitHub Pull Requests with proper structure and documentation. Invoke when user needs to create a PR for code changes or feature implementations."
---

# PR Creator (CI/CD)

## 功能描述 (Description)

This skill helps create well-structured GitHub Pull Requests with proper documentation, ensuring that CI/CD pipelines run smoothly and code reviews are efficient.

本技能帮助创建结构良好的 GitHub Pull Request，确保 CI/CD 管道顺利运行且代码审查高效。

## 核心功能 (Key Features)

- **PR 结构设计**：指导创建具有清晰标题、描述和标签的 PR
- **分支管理**：提供分支命名和管理的最佳实践
- **CI/CD 配置**：确保 PR 包含必要的 CI/CD 配置
- **代码审查准备**：帮助准备代码审查所需的文档和测试
- **变更说明**：指导编写清晰的变更说明和影响分析
- **合并策略**：提供 PR 合并的最佳实践

## 使用场景 (Usage Scenarios)

当用户需要以下帮助时，会触发此技能：
- 创建新功能或修复 bug 的 PR
- 准备代码审查的 PR
- 确保 PR 通过 CI/CD 检查
- 管理多个提交的 PR
- 处理复杂的代码变更

## 最佳实践 (Best Practices)

- 使用语义化的 PR 标题和描述
- 包含相关的 issue 链接
- 提供清晰的变更说明
- 确保所有测试通过
- 使用适当的标签和里程碑
- 遵循项目的贡献指南

## 参考 (Reference)

Based on the GitHub repository: https://github.com/google-gemini/gemini-cli/tree/main/.gemini/skills/pr-creator