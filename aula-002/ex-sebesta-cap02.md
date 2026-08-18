### Exercícios do Capítulo 02 do Sebesta

**1.** A genealogia das linguagens não é uma escada de progresso.
Explique essa afirmação e apresente dois fatores históricos que fazem
uma linguagem influenciar outra sem necessariamente substituí-la.

**R:** A frase em questão se refere ao desenvolvimento das linguagens de programação, onde, este, não ocorre em uma linha reta de evolução que apresente níveis superiores e/ou avançados em comparação com as linguagens anteriores.

**Fatores de influência:**

- Surgimento de novos hardwares/arquiteturas: Evolução e mudanças a nível de hardware, podem alterar os objetivos de projeto das linguagens. Ex: O computador IBM 704, que trazia instruções de indexação e ponto flutuante diretamente no hardware, tendo eliminado a necessidade de pseudocódigos e sistemas de interpretação pura, servindo de base para o nascimento do Fortran.
- Evolução de metodologias de projeto de software: Com o surgimento de novas formas de pensar na estruturação de código, houve demanda de novos recursos linguísticos. A transição da orientação a procedimentos para a orientação a dados, fez com que conceitos de classe e herança, introduzidos por linguagens como SIMULA 67, influenciassem o design de linguagens posteriores, como Java e C#.

**2.** Plankalkül não foi implementada em sua época. Ainda assim, por que
ela é relevante para a história das linguagens? Cite três recursos
antecipados por seu projeto e explique o valor de um deles.

**R:** Plankalkül é altamente relevante para a história das linguagens, pois continha programas de complexidade muito maior do que qualquer outro escrito antes de 1945. Tinha incluso, ordenação de vetores, testes de conectividade de grafos e era capaz de jogar xadrez.

**Três recursos:**

- Inclusão de vetores e structs aninhados.
- Sentenças iterativas estruturadas.
- Inclusão de asserções matemáticas.

Explicando asserções: As asserções matemáticas antecipadas por Plankalkül, possuem um grande valor histórico, pois estabeleceram as bases conceituais para a verificação de programas e para a programação defensiva. Com isso, foi possível criar restrições que garatem a estabilidade e a corretude lógico do código duranre execução.

**3.** Compare Short Code, Speedcoding e os sistemas A-0/A-1/A-2 quanto
ao problema enfrentado e à estratégia adotada. Por que chamá-los
simplesmente de compiladores modernos seria impreciso?

**R:**

- **Short Code:**
  - **Problema enfrentado:** Falta de linguagens de alto nível.
  - **Estratégia adotada:** Representação codificada de expressões matemáticas em palavras de memória de 72 bits.
- **Speedcoding:**
  - **Problema enfrentado:** Hardware escasso para operações de ponto flutante e indexação de vetores nos computadores.
  - **Estratégia adotada:** Conversão do computador em uma calculadora virtual de ponto flutuante e inclusão de um sistema de interpretação pura.
- **Sistemas A-0 / A-1 / A-2:**
  - **Problema enfrentado:** Falta de ferramentas automáticas para reutilizar rotinas comuns.
  - **Estratégia adotada:** Funcionamento similar ao de um macro-expansor em linguagem de montagem.

**Por que não chamá-los de compiladores modernos?:**
R: O termo se torna inadequado devido, Short Code e o Speedcoding serem baseados em interpretação pura, ou seja, estes não traduziam o código-fonte em um arquivo executável, para isso tinha um software que agia como uma maquina virtual. Já os sistemas A-0 / A-1 / A-2, funcionavam na epoca de forma a expandir pseudocódigos em sub-rotinas de código de máquina preexistentes, essa estratégia se assemelha muito mais a um macro-montador. Além disso por apresentarem um gargalo de desempenho que compiladores modernos não possuem.

**4.** Explique por que o projeto Fortran precisou convencer
programadores de que código traduzido podia competir com código de
máquina escrito à mão. Relacione desempenho, custo de programação
e adoção.

**R:** O projeto Fortran precisou realizar esse esforço devido aos seguintes fatores:

- **Desempenho:** Na epóca, as memórias eram minúsculas e o tempo de processamento era extremamente caro.
- **Custo de programação:** Programar diretamente em código de máquina manuscrito gerava alta eficiência de execução, mas era um processo lento e tedioso, o que elevava o custo do tempo gasto no desenvolvimento.
- **Adoção:** Para que os programadores aceitassem adotar a nova linguagem, a equipe do Fortran teve de focar intensamente em técnicas pioneiras de otimização de código.

**5.** Lisp surgiu em um contexto diferente de Fortran. Compare os
domínios, a representação de dados e o estilo de computação
favorecido pelas duas linguagens.

**R:**

- **Domínios de programação:** O Fortran foi projetado para aplicações para aplicações científicas de alta perfomance, enquanto o Lisp surgiu voltado para a área de Inteligência Artificial.
- **Representação de dados:** O Fortran utiliza estruturas de dados simples como vetores e matrizes numéricas. O Lisp, baseia-se em dados simbólicos estruturados em listas encadeadas.
- **Estilo de computação:** O Fortran adota o paradigma imperativo, procedural, ditado pelo estado de variáveis e iterações. O Lisp foi concebido como uma linguagem funcional, onde a computação ocorre ocorre pela aplicação de funções e uso de recursão.

**16.** Compare Perl, JavaScript, PHP, Python, Ruby e Lua usando três eixos: domínio inicial, estruturas de dados e estratégia de implementação. Evite concluir que todas são iguais por serem chamadas de scripting.

**R:** Apesar de frequentemente agrupadas sob o rótulo de "linguagens de scripting", essas linguagens possuem distinções fundamentais em seu design e propósito original.

- **Perl:**
  - **Domínio inicial:** Originalmente criada como um utilitário para processamento de arquivos de texto e administração de sistemas UNIX. Posteriormente ganhou força como linguagem CGI para a Web.
  - **Estruturas de dados:** Variáveis escalares, vetores de tamanho dinâmico (que podem ser esparsos) e matrizes associativas (hashes).
  - **Estratégia de implementação:** Sistema híbrido. Os programas são primeiro compilados para uma linguagem intermediária (para detecção de erros) antes de serem interpretados.

- **JavaScript:**
  - **Domínio inicial:** Programação no lado do cliente (client-side) para a Web, visando adicionar interatividade a documentos HTML estáticos.
  - **Estruturas de dados:** Arrays dinâmicos e objetos baseados em protótipos, que funcionam essencialmente como coleções dinâmicas de propriedades e métodos.
  - **Estratégia de implementação:** Originalmente usava interpretação pura nos navegadores, mas hoje utiliza amplamente a compilação JIT (Just-in-Time) para maior eficiência.

- **PHP:**
  - **Domínio inicial:** Scripts no lado do servidor (server-side) específicos para desenvolvimento Web, gerando código HTML dinâmico.
  - **Estruturas de dados:** Arrays nativos de PHP, que misturam características de vetores indexados tradicionais e matrizes associativas.
  - **Estratégia de implementação:** Inicialmente interpretação pura, mas evoluiu para sistemas de implementação híbridos.

- **Python:**
  - **Domínio inicial:** Administração de sistemas, computação científica e programação de propósito geral.
  - **Estruturas de dados:** Listas (vetores dinâmicos), tuplas (imutáveis) e dicionários (matrizes associativas).
  - **Estratégia de implementação:** Sistema híbrido, onde o código é compilado para bytecodes e então executado por uma máquina virtual.

- **Ruby:**
  - **Domínio inicial:** Scripting com um foco absoluto e purista no paradigma de Orientação a Objetos.
  - **Estruturas de dados:** Todos os dados são objetos, incluindo vetores e hashes (matrizes associativas) fortemente integrados à linguagem.
  - **Estratégia de implementação:** Historicamente interpretação pura, o que afetava sua eficiência, transitando para abordagens híbridas em versões mais modernas.

- **Lua:**
  - **Domínio inicial:** Scripting de extensão, desenvolvida para ser uma linguagem embarcada dentro de outras aplicações (muito comum na indústria de jogos).
  - **Estruturas de dados:** Baseia-se quase exclusivamente em tabelas, que são matrizes associativas flexíveis usadas para representar vetores, registros e objetos.
  - **Estratégia de implementação:** Sistema híbrido compilado para bytecodes extremamente leves e rápidos, executados em uma máquina virtual.

**17.** C# foi apresentada como evolução no ambiente .NET. Compare duas decisões de C# com suas correspondentes em Java ou C++ e explique o problema que pretendem resolver.

**R:** A linguagem C# foi construída a partir da fundação de C++ e Java, com o objetivo de corrigir falhas ou deficiências dessas linguagens dentro da plataforma .NET.

- **Decisão 1: A sentença de seleção múltipla (`switch`)**
  - **Correspondente em C++:** O `switch` de C e C++ permite o fall-through implícito. Ou seja, se o programador esquecer de colocar um `break` no final de um `case`, o fluxo de execução continuará executando os comandos do próximo `case`.
  - **O problema que C# resolve:** Esse comportamento no C++ frequentemente causa bugs difíceis de rastrear. Em C#, a regra de semântica estática foi alterada: todo segmento `case` (que não seja vazio) deve terminar explicitamente com uma instrução de desvio, como `break` ou `goto`. Isso resolve o problema de falta de confiabilidade gerado por esquecimentos acidentais.

- **Decisão 2: O tratamento de referências a métodos (Delegates)**
  - **Correspondente em C++:** Em C++, para referenciar subprogramas/métodos como parâmetros ou passá-los adiante, utilizam-se os ponteiros de função.
  - **O problema que C# resolve:** Ponteiros de função no C/C++ são notoriamente inseguros em relação aos tipos. C# introduziu o conceito de representantes (Delegates), que são objetos orientados a objetos e fortemente tipados que podem armazenar referências seguras para os métodos. Isso mantém o poder flexível de passagem de métodos sem perder a confiabilidade imposta pela verificação de tipos.

**18.** Diferencie XSLT e JSP quanto a entrada, processamento e saída. Por que ambas podem ser chamadas de linguagens híbridas de marcação e programação?

**R:**

- **XSLT (eXtensible Stylesheet Language Transformations):**
  - **Entrada:** Um documento estruturado em XML.
  - **Processamento:** Transforma e manipula os dados e a estrutura da árvore do XML de entrada através de regras e templates de formatação.
  - **Saída:** Um novo documento estruturado (que pode ser outro XML, ou frequentemente um HTML para apresentação).

- **JSP (JavaServer Pages):**
  - **Entrada:** Requisição de um cliente apontando para um documento JSP que contém a estrutura de apresentação.
  - **Processamento:** A execução (lado do servidor) dos blocos de código Java que estão embutidos ou mesclados ao longo do documento usando tags especiais ou bibliotecas (como JSTL).
  - **Saída:** Um documento de marcação (comumente HTML) gerado de forma dinâmica para exibição no navegador do usuário.

**Por que são linguagens híbridas?**
R: Ambas recebem essa denominação porque quebram o conceito tradicional de que marcação não é programação. Elas combinam uma base de linguagem de marcação (XML ou HTML) com estruturas de controle comuns à linguagens de programação imperativas, como instruções de seleção (ifs) e de repetição (laços de iteração), mesclando definição estrutural/apresentação com lógica computacional.

**19.** Crie uma linha do tempo com oito linguagens de pelo menos quatro paradigmas. Para cada ligação, escreva o tipo de influência; não use apenas setas cronológicas.

**R:** Abaixo segue uma linhagem cronológica ilustrando o processo evolutivo das linguagens, abrangendo os paradigmas Imperativo, Funcional, Lógico e Orientado a Objetos:

- **Fortran (1957) [Paradigma Imperativo]:** Primeira linguagem compilada de alto nível. Focou fortemente em performance e computação matemática usando vetores.
- **Lisp (1959) [Paradigma Funcional]:** Desenvolvida para necessidades de inteligência artificial. Influência: Afastou-se da base da arquitetura de von Neumann adotada no Fortran, introduzindo processamento de listas, tipagem dinâmica e suporte para recursão na computação.
- **ALGOL 60 (1960) [Paradigma Imperativo]:** Influência: Tomou as fundações imperativas de Fortran, mas universalizou e formalizou sua sintaxe. Introduziu a estrutura de blocos e passagem por valor/nome, influenciando severamente o design semântico e escopo de quase todas as linguagens imperativas subsequentes.
- **SIMULA 67 (1967) [Paradigma Orientado a Objetos (Iniciador)]:** Influência: Utilizou como base estrutural o ALGOL 60, estendendo-o fortemente para incorporar as bases da abstração de dados: classes, instâncias (objetos) e herança básica.
- **Prolog (1972) [Paradigma Lógico]:** Influência: Criado a partir do cálculo de predicados de primeira ordem. Diferenciou-se de Fortran, ALGOL e Lisp ao assumir uma abordagem baseada em regras declarativas onde o desenvolvedor define a base de dados de inferências, delegando o fluxo ao motor de inferência e ao rastreamento (backtracking) do sistema.
- **Smalltalk (1980) [Paradigma Orientado a Objetos (Puro)]:** Influência: Lapidou o conceito inicial do SIMULA 67, removendo características puramente procedurais e implementando um modelo purista onde absolutamente tudo na linguagem funciona e reage via troca de mensagens entre objetos.
- **C++ (1985) [Híbrido: Imperativo / Orientado a Objetos]:** Influência: Herdou as fundações de sintaxe e o uso em nível de sistema de sua base, a linguagem C, combinando ativamente aos conceitos criados pelo SIMULA 67 e solidificados pelo Smalltalk (classes e herança), permitindo o desenvolvimento com orientação a objetos aliado à altíssima eficiência em um mesmo ambiente.
- **Java (1995) [Paradigma Orientado a Objetos]:** Influência: Herdou de forma quase idêntica as estruturas sintáticas de controle de C++, mas agiu como uma resposta direta às inseguranças de C++. Removeu explicitamente conceitos falhos (como ponteiros diretos de memória e herança múltipla irrestrita) a favor de segurança (coleta de lixo automática) e confiabilidade estrita.

**20.** Estudo de caso: uma equipe precisa escolher tecnologias para cálculo científico, regras declarativas, aplicação Web interativa e firmware restrito. Proponha famílias de linguagens, justifique historicamente cada escolha e explicite dois trade-offs.

**R:** Baseado na relação histórica do design de linguagens com seus domínios pretendidos:

- **Cálculo Científico: Família Fortran (ex: Fortran 95/2003)**
  - **Justificativa Histórica:** O Fortran foi a primeira linguagem de alto nível, projetada especificamente na década de 1950 com o propósito de rodar eficientemente matrizes e fórmulas matemáticas simulando engenharia.
  - **Trade-offs:** Você obtém um altíssimo grau de eficiência na execução de ponto flutuante e paralelização; no entanto, em troca, obtém menor expressividade de código e piores capacidades na implementação de lógicas que fogem da matemática padrão.

- **Regras Declarativas: Família de Linguagens Lógicas (ex: Prolog)**
  - **Justificativa Histórica:** O Prolog originou-se na década de 1970 da necessidade acadêmica e de IA de trabalhar com dados simbólicos provando deduções com o cálculo de predicados.
  - **Trade-offs:** Você ganha uma facilidade de escrita enorme em problemas de banco de conhecimento e regras, conseguindo montar relações com poucas linhas; em troca, perde radicalmente em eficiência de custo computacional/tempo de execução e memória, devido à engine de resolução.

- **Aplicação Web Interativa: Família de Linguagens de Scripting / Dinâmicas (ex: JavaScript ou PHP)**
  - **Justificativa Histórica:** Desenvolvidas diretamente no final dos anos 1990 como uma ponte em navegadores (JavaScript) e processadores (PHP) para adicionar conteúdos vivos dinâmicos onde o HTML era incapaz de atuar nativamente.
  - **Trade-offs:** Elevada flexibilidade e rápida resposta na hora de criar interatividade de rede, porém, existe a perda de confiabilidade pelo fato dessas linguagens geralmente ignorarem certas detecções de erro através de coerções e tipagem dinâmica implícitas.

- **Firmware Restrito: Família Imperativa Orientada a Sistemas (ex: C)**
  - **Justificativa Histórica:** A linguagem C foi pensada primariamente como ferramenta para o desenvolvimento dos sistemas operacionais (UNIX) permitindo uma abstração da linguagem Assembly de baixo nível, mas com a mesma filosofia de controle e otimização.
  - **Trade-offs:** Total controle da capacidade alocativa da memória e um desempenho inigualável em hardware com pouco processamento (firmware), porém, troca essa excelência por um custo em segurança e confiabilidade (visto a ausência rigorosa de limitações em índices ou gerenciadores em ponteiros).
