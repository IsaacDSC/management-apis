# management-apis
##### Utilizar um serviço Backend for Frontend (BFF) para comunicar com múltiplos serviços backend pode trazer diversos benefícios, mas também apresenta alguns tradeoffs. Aqui estão alguns pontos a considerar:  

### Benefícios
Agregação de Dados: O BFF pode agregar dados de múltiplos serviços backend em uma única resposta, reduzindo a quantidade de chamadas de rede necessárias para o cliente.
Customização para Clientes: Diferentes clientes (web, mobile, etc.) podem ter diferentes necessidades de dados. Um BFF pode fornecer endpoints específicos para cada tipo de cliente, otimizando a performance e a experiência do usuário.
Segurança: O BFF pode atuar como uma camada de segurança adicional, gerenciando autenticação e autorização antes de encaminhar as solicitações para os serviços backend.
Simplificação do Cliente: O cliente pode ser simplificado, pois não precisa conhecer a estrutura dos serviços backend. Toda a lógica de comunicação e agregação de dados é gerenciada pelo BFF.
Modularidade: Facilita a modularidade dos serviços, permitindo que cada serviço backend evolua independentemente, enquanto o BFF gerencia a integração.

### Tradeoffs
Complexidade Adicional: Introduzir um BFF adiciona uma camada extra de complexidade ao sistema. O BFF precisa ser mantido e atualizado conforme os serviços backend mudam.
Ponto Único de Falha: O BFF pode se tornar um ponto único de falha. Se o BFF estiver indisponível, todos os clientes que dependem dele também serão afetados.
Latência: Cada chamada ao BFF pode introduzir latência adicional, especialmente se o BFF precisar fazer múltiplas chamadas aos serviços backend para compor a resposta.
Manutenção: O BFF precisa ser mantido em sincronia com os serviços backend. Mudanças na API dos serviços backend podem exigir atualizações no BFF.
Escalabilidade: O BFF precisa ser escalável para lidar com a carga de todas as solicitações dos clientes. Isso pode exigir recursos adicionais e planejamento de infraestrutura.****



### FALTA
- Logs
- Ambiente para envio para s3 e não guardar dados na aplicação
- Adapter do banco para suportar grandes cargas (DynamoDB)
- Config para trabalhar com async em casos de escritas
- Coverage
- Commands da aplicação
- Readme
- GraphQL