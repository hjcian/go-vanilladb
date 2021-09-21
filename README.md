# go-vanilladb
- [go-vanilladb](#go-vanilladb)
  - [Questions](#questions)
  - [學習素材與方法](#學習素材與方法)
  - [延伸閱讀](#延伸閱讀)

## Questions

有關 RDBMS，我很好奇：
- 如何實作出 [ACID](https://en.wikipedia.org/wiki/ACID) 特性？
- 如何將一筆 record 儲存到 file system？
- 如何索引資料? 索引又是如何儲存到 file system？
- 如何處理 crash 時、保證資料不遺失？
- 如何處理高併發情境？

## 學習素材與方法

- 清大課程-雲端資料庫
  - CS 471000 Introduction to Database Systems - Implementation, architectural design, and trade-offs. https://nthu-datalab.github.io/db/
  - [VanillaDB source code (JAVA)](https://github.com/vanilladb/vanillacore)
- 學習方法
  - 1) 看清大課程的 videos 及 lectures，了解基礎版的 RDBMS 可以怎麼實作，也就是了解課程中的 VanillaDB 的系統架構及實作原理
  - 2) 看懂 VanillaDB JAVA codes，並且以 Go 重刻/重構
  - 3) 分享筆記，凝聚共識，協作討論，共同貢獻程式碼


## 延伸閱讀
- Golang Learning |【Go 夜读】#117 详解 rosedb 及存储模型
  - https://www.youtube.com/watch?v=Knh7EYfVIKs
  - https://talkgo.org/t/topic/2386
