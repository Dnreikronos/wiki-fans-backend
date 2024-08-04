<h1>Wiki-Fans API</h1>

<p>This is an API developed in Go Lang to provide information to the Frontend being developed in the project <a href="https://github.com/Dnreikronos/wiki-fans-kuroko-no-basket">wiki-fans-kuroko-no-basket</a>.</p>


<h2>Table of Contents</h2>
    <ul>
        <li><a href="#introduction">Introduction</a></li>
        <li><a href="#features">Features</a></li>
        <li><a href="#api-endpoints">API Endpoints</a></li>
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
    
<h2 id="things-learned"> Concepts Learned doing this project</h2>
<p>Here are some concepts I learned developing this specific project:</p>
    <ul>
        <li><strong>How to manage environments variables;</li>
        <li><strong>How to use the Godotenv library</li>
        <li><strong>How to use Gitignore for hide some important informations</li>
        <li><strong>How to create a PostgreSQL database using Docker</li>                
    </ul>

<h2 id="api-endpoints">API Endpoints</h2>
    <p>Here are some of the main endpoints provided by the API:</p>
    <ul>
        <li><strong>GET /api/characters:</strong> Retrieves a list of characters.</li>
        <li><strong>GET /api/characters/{id}:</strong> Retrieves details of a specific character by ID.</li>
        <li><strong>POST /api/characters:</strong> Creates a new character (requires authentication).</li>
    </ul>
