// imports

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import BootstrapTable from 'reactjs-bootstrap-table';

// functional component
function Employee() {
    useEffect(() => {
        getEmployees()
    } , [])

    const [employeeData, setData] = useState([]);
    const url = 'http://localhost:8000/'


  const getEmployees = () => {
    axios.get(`${url}employees/`)
    .then(res => {
      const allEmployees = res.data
      setData(allEmployees.data) 
      console.log(allEmployees.data)
    }).catch(err => {
      console.log(err)
    })

  }

  const employeeColumns = [
    {
    name: "id",
    display: "ID"
    },
    {
    name: "first_name",
    display: "First Name"
    },
    {
    name: "last_name",
    display: "Last Name"
    },
    {
    name: "reports_to",
    display: "Reports To"
    },
    {
    name: "birth_date",
    display: "Birth Date"
    },
    {
    name: "hire_date",
    display: "Hire Date"
    },
    {
    name: "address",
    display: "Address"
    },
    {
    name: "city",
    display: "City"
    },

    {
    name: "state",
    display: "State"
    },

    {
    name: "country",
    display: "Country"
    },
    {
    name: "postal_code",
    display: "Postal Code"
    },
    {
    name: "phone",
    display: "Phone"
    },
    {
    name: "email",
    display: "Email"
    }
]
// jsx below

return (
    <div className="table-container">
       {<BootstrapTable tableClass={"table"} headers = {true} columns={employeeColumns} data={employeeData} />}

    </div>
    );
}

export default Employee;