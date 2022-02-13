// imports
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import BootstrapTable from 'reactjs-bootstrap-table';

// functional component
function Customer() {
    useEffect(() => {
        getCustomers()
    } , [])

    const [customerData, setData] = useState([]);
    const url = 'http://localhost:8000/'
  
    const getCustomers = () => {
        axios.get(`${url}customers/`)
        .then(res => {
          const allCustomers = res.data
          setData(allCustomers.data)
          console.log(allCustomers.data)
        }).catch(err => {
          console.log(err)
        })
    
      }

      console.log(customerData)

      const customerColumns = [
        {
        name: "id",
        display: "ID"
        },
        {
        name: "customer_first_name",
        display: "Customer First Name"
        },
        {
        name: "customer_last_name",
        display: "Customer Last Name"
        },
        {
        name: "customer_company",
        display: "Customer Company"
        },
        {
        name: "customer_email",
        display: "Customer Email"
        },
        {
        name: "customer_phone",
        display: "Customer Phone"
        },
        {
        name: "invoice_date",
        display: "Invoice Date"
        },
        {
        name: "billing_address",
        display: "Address"
        },

        {
        name: "billing_city",
        display: "City"
        },

        {
        name: "billing_country",
        display: "Country"
        },
        {
          name: "billing_postal_code",
          display: "Postal Code"
        },
        {
          name: "billing_state",
          display: "State"
        },
        {
          name: "unit_price",
          display: "Unit Price"
        },
        {
          name: "quantity",
          display: "Quantity"
        }
    ]
    //  jsx below
    return (
      <div className="table-container">
          {<BootstrapTable tableClass={"table"} headers = {true} columns={customerColumns} data={customerData} />}
      </div>
      );
  }
  
  export default Customer;
