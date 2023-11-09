# How does Browser work?

## Internet Protocols

### DNS

Domain Name System (DNS) is basically a phone book for the internet. When you type a web address in your browser, your computer uses DNS to look up the IP address of the website. 

DNS is a hierarchical and decentralized naming system for computers, services, or other resources connected to the Internet or a private network. It associates various information with domain names assigned to each of the participating entities.

DNS has multiple record types like A, AAAA, CNAME, MX, NS, PTR, SOA, SRV, TXT, etc. The most important one is A record which maps a domain name to an IP address.

### HTTP(S)

Hypertext Transfer Protocol (HTTP) is an application-layer protocol for transmitting hypermedia documents, such as HTML. It was designed for communication between web browsers and web servers, but it can also be used for other purposes. HTTP follows a classical client-server model, with a client opening a connection to make a request, then waiting until it receives a response. HTTP is a stateless protocol, meaning that the server does not keep any data (state) between two requests.

The 'S' in HTTPS stands for 'Secure'. Basically, an HTTP on top of a Secured Socket Layer (SSL) and Transport Layer Security (TLS). It means all communications between your browser and the website are encrypted. HTTPS is often used to protect highly confidential online transactions like online banking and online shopping order forms.

## Browser Policies

### Content Security Policy (CSP)

Content Security Policy (CSP) is an added layer of security that helps to detect and mitigate certain types of attacks, including Cross-Site Scripting (XSS) and data injection attacks. These attacks are used for everything from data theft, to site defacement, to malware distribution.


### Same-Origin Policy (SOP)

The same-origin policy is a critical security mechanism that restricts how a document or script loaded from one origin can interact with a resource from another origin. It helps isolate potentially malicious documents, reducing possible attack vectors.

### Cross-Origin Resource Sharing (CORS)

Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional HTTP headers to tell browsers to give a web application running at one origin, access to selected resources from a different origin. A web application executes a cross-origin HTTP request when it requests a resource that has a different origin (domain, protocol, or port) from its own.

# The Browser

## The Browser Engine

The browser engine marshals actions between the browser’s user interface and the browser’s rendering engine.

When you type in a new website and press the enter key, the browser UI will tell the browser engine, which will then communicate with the rendering engine.

## The Rendering Engine

The rendering engine is responsible for displaying the requested content.

The rendering engine will start by getting the contents of the requested document from the networking layer.

It takes in the HTML code and parses it to create the DOM (Document Object Model) tree.

The rendering engine will then parse the CSS to build the CSSOM (CSS Object Model). It’s like the DOM, but for the CSS rather than the HTML.

While the CSS is being parsed and the CSSOM is being created, the browser is downloading other assets through the Network Layer like JavaScript files.

The rendering engine communicates with the JavaScript engine to execute the JavaScript code and manipulate the DOM and CSSOM.

The rendering engine then takes the DOM and the CSSOM and combines them to create the Render tree.

The rendering engine then uses the UI backend for laying out the website on the screen and finally painting the pixels to the screen.

The entire process that the rendering engine goes through is called the Critical Rendering Path.

![](https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,format=auto,onerror=redirect,quality=80/uploads/asset/file/f004f954-de8a-473b-81f4-f1bde3750129/2dae3e28-cb30-4fcd-b477-8d22e87b9a32_880x254.jpeg)

## Networking Layer

The Networking Layer is responsible for making network calls to fetch resources.

It imposes the right connection limits, formats requests, deals with proxies, caching, and much more.

## JavaScript Engine

The JavaScript Engine is used to parse and execute JavaScript code on the DOM or CSSOM. The JavaScript code is provided by the web server, or it can be provided by the web browser (browser extensions or features of the browser like automatic ad-blocking).

Early browsers used JavaScript interpreters, but modern JavaScript engines use Just-In-Time compilation for improved performance.

## UI Backend

This layer is responsible for drawing the basic widgets like select or input boxes and windows. Underneath it uses operating system UI methods.

The rendering engine uses the UI backend layer during the layout and painting stages to display the web page on the browser.

# Data Storage

The browser needs to save data locally (cookies, cache, etc.) so the Data Storage component handles this part.

Modern browsers also support storage mechanisms like localStorage, IndexedDB, and FileSystem.

# Server-side Rendering vs Client-side Rendering

## Server-side Rendering

Server-side rendering (SSR) is the process of rendering web pages on a server and passing them to the browser (client-side), instead of rendering them in the browser.

The main advantage of SSR is that the user will see the content faster. The browser doesn’t have to wait for all the JavaScript to be downloaded and executed before displaying the content.

The main disadvantage of SSR is that it’s slower to render the initial page. The server has to do all the work of rendering the page, which can take a lot of time. It’s also more difficult to implement.

## Client-side Rendering

Client-side rendering (CSR) is the process of rendering web pages in the browser using JavaScript.

The main advantage of CSR is that it’s faster to render the initial page. The server doesn’t have to do all the work of rendering the page, which can take a lot of time. It’s also easier to implement.

The main disadvantage of CSR is that the user will see the content slower. The browser has to wait for all the JavaScript to be downloaded and executed before displaying the content.
