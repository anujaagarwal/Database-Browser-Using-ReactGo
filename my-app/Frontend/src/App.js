// imports
import './App.css';
import React, { useState, useEffect } from 'react';
import Actor from './Actor';
import Artist from './Artist';
import Customer from './Customer';
import Employee from './Employee';
require('bootstrap/dist/css/bootstrap.css');



// App function
function App() {
  return (
    <div className="App">
      <p>Actor Table</p>
      <Actor />
      <p>Artist Table</p>
      <Artist />
      <p>Customer Table</p>
      <Customer />
      <p>Employee Table</p>
      <Employee />

    </div>
  );
}

export default App;
