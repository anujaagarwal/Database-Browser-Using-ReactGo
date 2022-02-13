// imports
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import BootstrapTable from 'reactjs-bootstrap-table';



// functional component
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

// jsx below
    return (
    <div className="Actor">
        {<BootstrapTable tableClass={"table"} resize={{extra: 200}} headers = {true} columns={actorColumns} data={actorData} />}
    </div>
    );
}

export default Actor;
