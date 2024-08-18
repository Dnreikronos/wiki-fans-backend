<h1>Wiki-Fans API</h1>

<p>This is an API developed in Go Lang to provide information to the Frontend being developed in the project <a href="https://github.com/Dnreikronos/wiki-fans-kuroko-no-basket">wiki-fans-kuroko-no-basket</a>.</p>

<h2>Table of Contents</h2>
<ul>
    <li><a href="#introduction">Introduction</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#api-endpoints">API Endpoints</a></li>
    <li><a href="#concepts-learned">Concepts Learned</a></li>
    <li><a href="#deployment">Deployment</a></li>
</ul>

<h2 id="introduction">Introduction</h2>
<p>The Wiki-Fans API is designed to serve as the backend for the wiki-fans-kuroko-no-basket frontend project. It provides basic comprehensive data and functionality to support a rich user experience for fans of the series Kuroko no Basket.</p>

<h2 id="features">Features</h2>
<ul>
    <li><strong>RESTful API:</strong> Provides a set of endpoints to interact with the data.</li>
    <li><strong>Chi Router:</strong> Utilizes the Chi router for fast and efficient routing.</li>
    <li><strong>TOML Configuration:</strong> Supports configuration via TOML files for easy setup and customization.</li>
    <li><strong>Scalable:</strong> Built with Go for performance and scalability.</li>
</ul>

<h2 id="concepts-learned">Concepts Learned</h2>
<p>Here are some concepts I learned developing this specific project:</p>
<ul>
    <li><strong>How to manage environment variables</strong></li>
    <li><strong>How to use the Godotenv library</strong></li>
    <li><strong>How to use Gitignore to hide some important information</strong></li>
    <li><strong>How to create a PostgreSQL database using Docker</strong></li>
    <li><strong>How to create, configure and mannage a EC2 instance on AWS</strong></li>
</ul>

<h2 id="api-endpoints">API Endpoints</h2>
<p>Here are some of the main endpoints provided by the API:</p>
<ul>
    <li><strong>GET /api/characters:</strong> Retrieves a list of characters.</li>
    <li><strong>GET /api/characters/{id}:</strong> Retrieves details of a specific character by ID.</li>
</ul>

<h2 id="deployment">Deployment</h2>
<p>Follow these steps to clone the repository and run the application on your machine:</p>
<ol>
    <li><strong>Clone the Repository:</strong>
        <pre><code>git clone https://github.com/Dnreikronos/wiki-fans-backend.git</code></pre>
    </li>
    <li><strong>Navigate to the Project Directory:</strong>
        <pre><code>cd wiki-fans-backend</code></pre>
    </li>
    <li><strong>Install Dependencies:</strong>
        <p>Ensure you have Go installed. Then, run:</p>
        <pre><code>go mod tidy</code></pre>
    </li>
    <li><strong>Run the Application:</strong>
        <p>Start the API server using:</p>
        <pre><code>go run main.go</code></pre>
    </li>
    <li><strong>Access the API:</strong>
        <p>The API will be accessible at <code>http://localhost:9000</code> by default.</p>
    </li>
    <li><strong>Optional - Docker:</strong>
        <p>If you prefer to run the application using Docker, ensure you have Docker installed, and then run:</p>
        <pre><code>docker-compose up --build</code></pre>
        <p>This will build and run the application in a Docker container.</p>
    </li>
</ol>

<p>You have now successfully set up and run the Wiki-Fans API on your local machine or server!</p>
