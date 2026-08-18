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
