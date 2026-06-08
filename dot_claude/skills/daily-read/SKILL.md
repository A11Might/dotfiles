---
name: daily-read
description: Use when wanting to read a random bookmark from Readeck, discover saved articles, or clear the reading backlog. Triggers include requests to "recommend an article", "random bookmark", "daily read", or "something to read".
---

# Daily Read

从 Readeck 书签库中随机推荐一篇未归档文章，由 Claude 阅读全文后生成中文摘要，帮助快速判断是否值得阅读。

## Prerequisites

- 环境变量 `READECK_API_TOKEN` 和 `READECK_API_URL` 已设置

## Steps

### 1. 获取随机书签

运行 `go run ~/.claude/skills/daily-read/scripts/main.go`，输出格式如下：

```
ID: xxx
标题: xxx
URL: xxx
来源: xxx
预计阅读: x 分钟
标签: xxx
描述: xxx
---
{文章全文}
```

如果文章内容为空或无意义（如登录页面），直接运行 `go run ~/.claude/skills/daily-read/scripts/main.go skip {id}` 跳过，然后重新获取。

### 2. 生成摘要

阅读 `---` 之后的完整文章内容，用中文写一段 **200 字以内**的摘要，突出：
- 文章的核心主题
- 关键观点或发现
- 帮助读者判断是否感兴趣的信息

### 3. 展示推荐

```
━━━━━━━━━━━━━━━━━━━━━━━━━━
📖 今日推荐阅读
━━━━━━━━━━━━━━━━━━━━━━━━━━
📌 标题: {title}
🌐 来源: {site}
🔗 链接: {url}
⏱  预计阅读: {x} 分钟

📝 AI 摘要:
{摘要}
━━━━━━━━━━━━━━━━━━━━━━━━━━
```

### 4. 询问用户操作

使用 AskUserQuestion 工具提供以下选项（多选关闭），根据用户选择执行对应操作：

- **打开链接** (`open {url}`) — 在浏览器中打开 Readeck 书签页，选择后继续询问下一步操作
- **下一篇** — 重新运行 `go run ~/.claude/skills/daily-read/scripts/main.go` 随机推荐下一篇
- **不感兴趣** — 运行 `go run ~/.claude/skills/daily-read/scripts/main.go skip {id}` 打标签 + 归档，然后自动推荐下一篇
- **归档** — 运行 `go run ~/.claude/skills/daily-read/scripts/main.go archive {id}` 直接归档，然后继续询问
- **退出** — 结束

### 5. 循环推荐

用户选择"下一篇"或"不感兴趣"后，自动回到步骤 1 继续推荐，直到用户选择退出或提示没有更多书签。
