# React

**React** is a Framework for Frontend Development. It is maintained by Facebook and a community of individual developers and companies.

It has a large ecosystem of tools and libraries that make it easy to build complex user interfaces.

# Before React?

A Web can be built by just HTML, CSS and JavaScript. But it is not easy to maintain and scale.

- HTML and CSS are not easy to reuse. JavaScript is not easy to maintain.

- Doing animation, interactive UI, data fetching, etc. are not easy.

- It is not easy to test.

- It is not easy to share code between platforms

# Module Bundler

A **Module Bundler** is a tool that takes pieces of JavaScript and their dependencies and bundles them into a single file, usually for use in the browser.

As from previous post about browser. When browser meets a `<script>` tag, it will halt the HTML parsing, fetch the JavaScript file, execute it, and then continue parsing the HTML.

Importing a JavaScript file from another JavaScript file is risky because the position of the `<script>` tags matters. A module bundler will take care of this problem by bundling all the JavaScript files into a single file.

A big drawback of a website is also that it is open to the public. Anyone can see the source code of a website. This is not good for security reasons. A module bundler will also take care of this problem by bundling all the JavaScript and making them practically unreadable by humans.

# React Component

A **React Component** is a JavaScript function that returns a React element.

A React element is a JavaScript object that represents a DOM element.

# React Hooks

**React Hooks** are functions that let you use React features in functional components.

## useState 

**useState** is a React Hook that lets you add state to functional components.

## useEffect

**useEffect** is a React Hook that lets you perform side effects in functional components.

## useContext

**useContext** is a React Hook that lets you use React Context in functional components.

## useRef

**useRef** is a React Hook that lets you create a mutable object.

# React Context

**React Context** is a way to pass data through the component tree without having to pass props down manually at every level.

# React Router

**React Router** is a library that allows you to handle routes in a web app.

# Virtual DOM

**Virtual DOM** is a JavaScript object that represents the DOM. It is a copy of the DOM.

When the state of the application changes, the Virtual DOM will be updated. Then, the Virtual DOM will be compared with the previous Virtual DOM. This is called **Diffing**.

The difference between the two Virtual DOMs will be used to update the DOM. This process is called **Reconciliation**.

Updating the actual DOM will cause the browser to re-render the page. This is an expensive operation. Virtual DOM allows React to update the DOM efficiently by only updating the parts of the DOM that have changed.

# JSX

**JSX** is a syntax extension to JavaScript. It is a JavaScript syntax that looks like HTML.

JSX is not a requirement for using React. But it is recommended to use JSX because it is easier to read and write.

JSX consists of HTML tags and JavaScript expressions. HTML tags are used to create React elements. JavaScript expressions are used to create dynamic content. JSX also has component tags that are used to create React components.

# Re-rendering

React will re-render a component when its state or props change. If a component is re-rendered, all its children will also be re-rendered. This can cause performance issues if the component tree is large.

Designing the component tree is an important part of React development. It is important to keep the component tree as small as possible. 

# High Order Component

**High Order Component** is a function that takes a component and returns a new component. It is used to share logic between components. It is also used to create reusable components. It is also used to create components that can be used in different contexts. It is also used to create components that can be used in different environments. 
