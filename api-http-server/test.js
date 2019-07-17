// test script for HTTP server

const path = require('path');
const fetch = require('node-fetch');

fetch('http://155.138.134.91/createAccount', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        id: 'jingleman',
        name: 'Jingle Man',
        email: 'jingle@sipher.co',
        accountType: 'Personal'
    })
});

let url = 'http://155.138.134.91/queryAccountByID?idToQuery=jingleman'
fetch(url, {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json',
    }
});