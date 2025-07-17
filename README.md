# github-explorer-cli

_Автор: [MichurinDev](https://github.com/MichurinDev)_

> Простая CLI-утилита на Go для поиска и просмотра быстрорастущих репозиториев GitHub по языку программирования и периоду (день, неделя, месяц).

## Описание

**github-explorer-cli** — небольшая кроссплатформенная команда, которая помогает быстро находить самые трендовые по количеству ⭐ репозитории на GitHub за день, неделю или месяц в интересующей вас технологии.

- Использует публичный GitHub API
- Позволяет фильтровать по языку программирования
- Удобный и лаконичный формат вывода
- Только стандартная библиотека Go

## Пример использования

Найти топ-10 репозиториев по Go за неделю
``github-explorer --lang python --period day --top 3``

_Вывод:_

```text
 1. Doriandarko/make-it-heavy ⭐ 679  Issues: 9
        A Python framework that emulates Grok Heavy functionality using intelligent multi-agent orchestration. Deploy 4 (or more) specialized AI agents in parallel to deliver comprehensive, multi-perspective analysis on any query.
        URL: https://github.com/Doriandarko/make-it-heavy

 2. chiphuyen/sniffly ⭐ 278  Issues: 2
        Claude Code dashboard with usage stats, error analysis, and sharable feature
        URL: https://github.com/chiphuyen/sniffly

 3. wzzheng/StreamVGGT ⭐ 229  Issues: 4
        Code for Streaming 4D Visual Geometry Transformer
        URL: https://github.com/wzzheng/StreamVGGT
```

## Аргументы командной строки

| Флаг         | Описание                                        | Значение по умолчанию |
|--------------|-------------------------------------------------|----------------------|
| `--lang`     | Язык программирования для поиска                | go                   |
| `--period`   | Период тренда: day, week, month                 | week                 |
| `--top`      | Количество репозиториев в списке                | 10                   |
| `-h, --help` | Показать справку                                |                      |

## Структура проекта

```github-explorer-cli/
├── main.go
├── internal/
│ └── app/
│ └── trending.go
├── go.mod
└── README.md
```
