# 基于LLM Agent监控产品演进探索

## LLM简介
LLM可以理解为小学生的人脑，具备有限的能力。使用LLM有一定的门槛，需要辅助一些技术手段才能更好的达到效果。

### LLM是怎么训练出来的
人所具备的能力是从出生后持续不断的学习、改进、迭代获得的。同样，LLM也是采用类似的方式获得智力的。LLM基于神经网络、概率论和训练数据，自身不断的学习、改进、迭代实现进化。

### 为什么无法评估LLM是如何工作的
LLM之前的软件都是人类按照确定的规则或方案构建的，所以人是可以解释这些软件的工作流程。而LLM是在没有人类协助的漫长迭代过程计算出来的，人只是给与了LLM一个大的图书馆，人很难完全了解LLM的思维过程，所以即使LLM的训练者也无法明确的回答LLM是如何工作的。正如一个人是无法完全理解另外一个人的脑回路是什么样子，只能猜测大致的思维过程。

### LLM语言
人类之间的交流是通过自然语言，自然语言的最小单位是词汇。同样，LLM有自己的语言，而语言的最小单位是token，等价于人类的词汇。人类有自己的词汇表，同样，LLM也有自己的由token构成的词汇表。token是唯一标识符，是一个数字。token对应自然语言的字符序列，也就是说，一个token可以比自然语言的词汇长或短，不是一一匹配的

### LLM是如何工作的 
LLM并没有想象中的智能，能够做的，只是猜测下一个词(token)是什么。工作过程大体如下：
1. 接收自然语言输入，转为token序列
2. 猜测(计算)下一个token
3. 基于新的token序列，不断迭代上述过程
4. 迭代完成后的token序列转为自然语言

### LLM的应用 
#### 常见用途
1. 自动化和效率。比如客户支持、数据分析、内容生成，自动化可以降低运维成本
2. 生成见解。比如专家服务
3. 创造更出色的客户体验。比如向客户提供高度个性化的内容，加强客户互动并改善用户体验
#### 挑战和局限
1. 成本高
2. 隐私和安全问题
3. 准确性和偏见

### 如何更好的使用LLM
因为LLM本身能力有限，为了更好的使用LLM，衍生出了很多技术手段来让LLM更好的满足用户的需求。本文主要介绍两个技术：
1. Prompt工程：与LLM交流的方法论
2. CoT(思维链)：一定程度补足LLM的推理能力

#### Prompt工程 
##### 什么是Prompt(提示词) 
Prompt是用户与语言模型(不仅仅大模型)的沟通方式，是用户给语言模型下达的指令输入(比如文本输入)，它包含以下任意要素：
1. 指令：想要模型执行的特定任务或者指令。
2. 上下文：引导大模型更好的响应。
3. 输入数据：用户输入的内容
4. 输出提示：指定输出类型或格式

##### 什么是Prompt工程 
提示工程是开发和优化提示词，因为大模型没有想象中的智能，Prompt工程可以让大模型更好的理解用户的需求。

#### CoT(思维链)
CoT是一种改进的Prompt技术，在特定的领域提高了大模型的可解释性、可信度和可控性：
1. 增强大模型的推理能力：帮助LLM更好的问题分解
2. 增强大模型的可解释性和可信度：LLM的解题过程可知
3. 增强大模型的可控性：LLM的解题过程可控
4. 增强大模型的灵活性和创造性：LLM不局限于“语言智能”

##### 缺点
1. 对用户有一定的要求：需要输入一些自然语言形式的推理过程
2. 不适用于专业知识或常识的问题：LLM本身没有掌握这些知识
3. 解决方案不够发散(创新)：确定的解题过程

##### CoT构造方式
1. 人工构造：质量高，但人力成本大，不好优化、不好跨任务迁移
2. 自动构造：长远目标。不稳定，幻觉问题

### 大、小模型的看法
大模型会越来越大直到继续增大模型无法带来收益。
苹果等手机厂商一定会搞出来手机上可以用的小模型。小模型可能会出现在各种领域，各种小模型。

## LLM Agent简介
LLM Agent的全称是LLM-Powered autonomous Agent(大模型驱动的具备自主能力的智能体)。

### 为什么会出现LLM Agent及现状
1. LLM更像是一个技术或者工具，但是人类需要的是一个产品或者解决方案，LLM Agent就是相应的解决方案。
2. LLM能力有限，需要针对不同场景或领域创造不同的Agent智能体来解决实际问题。

LLM Agent处于在爆发的初期，有非常多的Agent项目，但是目前还没有出现优秀的Agent范本。

### LLM Agent构成 
LLM Agent = LLM + Profile + Planning + Memory + Tool

#### Profile(角色)
Profile解决的问题是告诉Agent他的角色，即告诉Agent需要解决的问题的背景信息。

#### Planning(规划)
Planning是最体现智能智力的模块：根据任务设定具体的执行方案，同时不断迭代改进，弥补LLM推理能力的不足。
1. 子目标分解：Agent将大任务拆分为更小的可管理的子目标，进而有效处理复杂任务。这里可以采用CoT或者ToT技术来实现。
2. 反思和完善：迭代改进。

#### Memory(记忆)
Memory主要解决的是行业知识的传递问题，可以让Agent拥有长期记忆和短期记忆，让他更智能(LLM不具备长期记忆能力)。

#### Tool(工具) 
Tool可以理解为人类的手，通过Tool执行不同的行为，主要是调用外部API。

## 基于LLM Agent演进监控产品
LLM Agent出来之后，有一切SaaS产品值得基于LLM Agent重构的观点。针对监控产品来说，LLM Agent能够提升监控产品的产品力，甚至可以重新定义智能监控，值得一定程度的探索。

下一代监控产品可能是SRE机器人，主要具备一下能力：
1. 提出稳定性专家建议：站点分析、建议
2. 执行用户提出的任务
3. 故障自动定位、处理
4. 报告自动生成
5. ...

### 产品力
1. 提供个性化的用户体验。不同客户、不同角色对监控关注的方面不一样，加强产品与客户的互动
2. 降低产品的使用成本：机器人对话交互即可
3. 提升客户的工作效率：智能化、自动化完成客户需求
4. 提供专家服务
5. 基于数据迭代产品：客户的自然语言数据就是最好的数据

### 现实问题 
#### 盈利
很难说引入LLM Agent后就能够提高盈利。客户不一定会之买单：
1. 没有颠覆性的能力
2. 性价比低
3. 中国人工成本比国外低
#### LLM选择
LLM使用哪个，有没有相应的资源可以利用

### 策略
1. 公有云探索LLM Agent，提高监控产品在公有云上的竞争力。专有云暂不考虑
2. 前期限定客户范围、限定使用范围
3. 针对具体场景落地：比如减低使用成本、提供个性化体验
4. 传统交互为主，自然语言为辅；基于客户输入数据调整产品方向
4. 一定额度的免费试用，超出额度计费

## 参考

* [LLM Agent现状和一些思考](https://zhuanlan.zhihu.com/p/679032270)
* [重新思考LLMs和Agents](https://juejin.cn/post/7312243176834809908)
* [LLM下半场之Agent](https://blog.csdn.net/buptgshengod/article/details/133577947)
* [探索未知：LLM Agent 应用开发的全新时代(视频)](https://www.bilibili.com/video/BV1W84y197qq/?vd_source=b33cd5d48827a412ff3a0592910b9ccc)
* [探索未知：LLM Agent 应用开发的全新时代(文字)](https://zhuanlan.zhihu.com/p/678807063)
* [LLM Agents现状和未来](https://www.wehelpwin.com/article/5173)
* [AI Agent 2024 ToB破局点](https://36kr.com/p/2577415992516485)
* [LLM Agent开发指南](https://www.waytoagi.com/question/10566)
* [AI Agent深度讲解](https://zhuanlan.zhihu.com/p/676544930)
* [大模型Prompt工程](https://cloud.tencent.com/developer/article/2361298)
* [大模型Prompt工程](https://blog.csdn.net/qq128252/article/details/133775725)
* [大模型Prompt工程](https://m.huxiu.com/article/2058854.html)
* [思维链开山之作](https://zhuanlan.zhihu.com/p/612136862)
* [思维链](https://blog.csdn.net/Julialove102123/article/details/135499567)
* [Agent能力解密](https://blog.csdn.net/youyi300200/article/details/132864191)
* [Agent](https://www.cnblogs.com/tgzhu/p/18144366)
* [直觉理解LLM是如何工作的](https://www.53ai.com/news/qianyanjishu/2025.html)
* [什么是大模型](https://www.redhat.com/zh/topics/ai/what-are-large-language-models)
* [Generative Agents: Interactive Simulacra of Human Behavior](https://dl.acm.org/doi/pdf/10.1145/3586183.3606763)
* [LLM-Agent-Paper-List](https://github.com/WooooDyy/LLM-Agent-Paper-List)
* [LLM Powered Autonomous Agents](https://lilianweng.github.io/posts/2023-06-23-agent/)
