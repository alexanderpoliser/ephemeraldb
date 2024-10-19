# NoSQL Database in Go  

This project provides a simple NoSQL database implementation in Go. It offers basic functionalities such as storing, retrieving, deleting data, and persisting buckets in JSON files. Additionally, the interactive menu allows users to manage data directly through the terminal.

---

## Table of Contents
- [Features](#features)  
- [Requirements](#requirements)  
- [Installation and Execution](#installation-and-execution)  
- [How to Use](#how-to-use)  
- [Project Structure](#project-structure)  
- [Contributing](#contributing)  
- [License](#license)

---

## Features
- **Add and update data:** Inserts or updates a value associated with a key in a bucket.
- **Retrieve data:** Retrieves specific data or lists all data in a bucket.
- **Delete bucket:** Completely removes a bucket and its contents.
- **Save to file:** Saves the contents of a bucket to a JSON file.
- **Load from file:** Loads data from a JSON file into a bucket.
- **List buckets:** Lists all available buckets with an estimated data size.

---

## Requirements
- Go 1.18 or higher  
- OS with terminal support (Windows, macOS, or Linux)  

---

## Installation and Execution

1. **Clone the repository:**  
   ```bash
   git clone https://github.com/alexanderpoliser/ephemeraldb.git
   cd ephemeraldb
   ```

2. **In your terminal run the program:**  
   ```bash
   go run cmd/main.go
   ```
---

## How to Use

After launching the program, you'll see a menu with the following options:  

1. **Set (Add/Update):** Adds or updates a value in a bucket.  
2. **Get (Retrieve):** Retrieves a specific key or lists all data in a bucket.  
3. **Delete Bucket:** Deletes a bucket and its contents.  
4. **Save:** Saves a bucket to a JSON file.  
5. **Load:** Loads data from a JSON file into a bucket.  
6. **List Buckets:** Lists all available buckets with approximate size.  
7. **Exit:** Exits the program.

### Example Usage

- **Add data:**  
  - Select option 1 from the menu.
  - Provide the bucket name and key.
  - Enter a JSON-formatted value.

  Example JSON input:
  ```json
  {
    "name": "Product A",
    "price": 49.99,
    "stock": 100
  }
  ```

- **Save bucket to file:**  
  - Select option 4.
  - Provide the bucket name and the desired filename (seting `.json` as the file format).
  - The bucket will be saved in the `data/` directory.

- **Load bucket from file:**  
  - Select option 5.
  - Provide the bucket name and the filename (seting `.json` as the file format) to load from.

---

## Project Structure

```
ephemeraldb/
├── cmd/                     # Main entry point of the application
│   └── main.go              # Code that starts the program
├── internal/                # Internal logic of the application
│   ├── db/                  # Database logic
│   │   └── db.go            # Implementation of the NoSQLDB
│   ├── menu/                # Logic for the interactive menu
│   │   └── menu.go          # Menu code
│   └── utils/               # Utility functions
│       └── utils.go         # Various utilities
├── data/                    # Directory to store JSON files
├── go.mod                   # Go module file
├── go.sum                   # Checksum file for dependencies
├── logo.txt                 # ASCII art displayed in the menu
└── README.md                # Project documentation
```

---

## Contributing

Feel free to contribute to this project! Fork the repository, create a new branch, and submit a pull request.

1. Fork the project  
2. Create a new branch: `git checkout -b feature/new-feature`  
3. Commit your changes: `git commit -m 'Add new feature'`  
4. Push to the branch: `git push origin feature/new-feature`  
5. Open a pull request  

---

## License

This project is distributed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Contact

If you have any questions or suggestions, feel free to reach out!

---

---

# Banco de Dados NoSQL em Go  

Este projeto oferece uma implementação simples de banco de dados NoSQL em Go. Ele permite armazenar, recuperar e excluir dados, além de salvar buckets em arquivos JSON. O menu interativo permite que o usuário gerencie dados diretamente pelo terminal.

---

## Índice
- [Funcionalidades](#funcionalidades)  
- [Requisitos](#requisitos)  
- [Instalação e Execução](#instalação-e-execução)  
- [Como Usar](#como-usar)  
- [Estrutura do Projeto](#estrutura-do-projeto)  
- [Contribuindo](#contribuindo)  
- [Licença](#licença)

---

## Funcionalidades
- **Adicionar e atualizar dados:** Insere ou atualiza um valor associado a uma chave em um bucket.
- **Recuperar dados:** Recupera um valor específico ou lista todos os dados de um bucket.
- **Excluir bucket:** Remove completamente um bucket e seus dados.
- **Salvar em arquivo:** Salva o conteúdo de um bucket em um arquivo JSON.
- **Carregar de arquivo:** Carrega dados de um arquivo JSON para um bucket.
- **Listar buckets:** Lista todos os buckets disponíveis com estimativa de tamanho.

---

## Requisitos
- Go 1.18 ou superior  
- Sistema operacional com suporte ao terminal (Windows, macOS ou Linux)  

---

## Instalação e Execução

1. **Clone o repositório:**  
   ```bash
   git clone https://github.com/alexanderpoliser/ephemeraldb.git
   cd ephemeraldb
   ```

2. **No terminal execute o programa:**  
   ```bash
   go run cmd/main.go
   ```
---

## Como Usar

Ao executar o programa, você verá um menu com as seguintes opções:  

1. **Set (Add/Update):** Adiciona ou atualiza um valor em um bucket.  
2. **Get (Retrieve):** Recupera uma chave específica ou lista todos os dados de um bucket.  
3. **Delete Bucket:** Remove um bucket e seus dados.  
4. **Save:** Salva um bucket em um arquivo JSON.  
5. **Load:** Carrega dados de um arquivo JSON para um bucket.  
6. **List Buckets:** Lista todos os buckets disponíveis com estimativa de tamanho.  
7. **Exit:** Encerra a aplicação.

### Exemplo de Uso

- **Adicionar dados:**  
  - Escolha a opção 1 no menu.
  - Informe o nome do bucket e a chave.
  - Insira um valor no formato JSON.

  Exemplo de valor JSON:
  ```json
  {
    "name": "Produto A",
    "price": 49.99,
    "stock": 100
  }
  ```

- **Salvar bucket em arquivo:**  
  - Escolha a opção 4.
  - Informe o nome do bucket e o nome do arquivo (colocando `.json` como formato do arquivo).
  - O bucket será salvo no diretório `data/`.

- **Carregar bucket de arquivo:**  
  - Escolha a opção 5.
  - Informe o nome do bucket e o nome do arquivo para carregar.

---

## Estrutura do Projeto

```
ephemeraldb/
├── cmd/                     # Entrada principal da aplicação
│   └── main.go              # Código que inicia o programa
├── internal/                # Lógica interna do aplicativo
│   ├── db/                  # Lógica do banco de dados
│   │   └── db.go            # Implementação da NoSQLDB
│   ├── menu/                # Lógica do menu interativo
│   │   └── menu.go          # Código do menu
│   └── utils/               # Funções auxiliares
│       └── utils.go         # Utilitários diversos
├── data/                    # Diretório para armazenar arquivos JSON
├── go.mod                   # Arquivo de módulo do Go
├── go.sum                   # Arquivo de checksum das dependências
├── logo.txt                 # Arte ASCII exibida no menu
└── README.md                # Documentação do projeto
```

---

## Contribuindo

Sinta-se à vontade para contribuir com este projeto! Faça um fork do repositório, crie uma branch e envie um pull request.

1. Faça um fork do projeto  
2. Crie uma nova branch: `git checkout -b feature/nova-funcionalidade`  
3. Commit suas mudanças: `git commit -m 'Adiciona nova funcionalidade'`  
4. Envie para a branch: `git push origin feature/nova-funcionalidade`  
5. Abra um pull request  

---

## Licença

Este projeto é distribuído sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## Contato

Caso tenha alguma dúvida ou sugestão, sinta-se à vontade para entrar em contato!

---