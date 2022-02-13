import React, { useState, useEffect } from 'react';
import axios from 'axios';
import BootstrapTable from 'reactjs-bootstrap-table';




function Actor() {
    useEffect(() => {
        getActors()
    } , [])
    const url = 'http://localhost:8000/'
    const [actorData, setData] = useState([]);
    const getActors = () => {
        axios.get(`${url}actors/`)
        .then(res => {
        const allActors = res.data
        console.log(allActors.data)
        setData(allActors.data)
        console.log(allActors)
        }).catch(err => {
        console.log(err)
        })
    }



let data = [
   { id: 1, 'firstName': '...', lastName: '...', address: '...'},
   { id: 2, 'firstName': '...', lastName: '...', address: '...'},
]



let columns = [
    {name: 'firstName' },
    {name: 'lastName' },
    {name: 'address' }
  ]
  

    const actorColumns = [
        {
        name: "id",
        display: "ID"
        },
        {
        name: "actor_first_name",
        display:"Actor First Name"
        },
        {
        name: "actor_last_name",
        display:"Actor Last Name"
        },
        {
        name: "film_id",
        display: "Film Id"
        },
        {
        name: "film_title",
        display: "Film Title"
        },
        {
        name: "film_description",
        display: "Film Description"
        },
        {
        name: "film_language_id",
        display: "Film Language Id"
        },
        {
        name: "film_release_year",
        display: "Film Release Year"
        },

        {
        name: "film_rating",
        display: "Film Rating"
        },

        {
        name: "category_name",
        display: "Category Name"
        }
    ]





    return (
    <div className="Actor">
        {<BootstrapTable tableClass={"table"} resize={{extra: 200}} headers = {true} columns={actorColumns} data={actorData} />}
    </div>
    );
}

export default Actor;
