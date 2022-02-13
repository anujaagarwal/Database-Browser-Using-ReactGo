// imports
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import BootstrapTable from 'reactjs-bootstrap-table';


// functional component
function Artist() {
    useEffect(() => {
        getArtists()
    } , [])
    const url = 'http://localhost:8000/'
  
    const [artistData, setData] = useState([]);
  
    const getArtists = () => {
      axios.get(`${url}artists/`)
      .then(res => {
        const allArtists = res.data
        setData(allArtists.data)
        console.log(allArtists.data)
      }).catch(err => {
        console.log(err)
      })
  
    }
    
    const artistColumns = [
        {
        name: "id",
        display: "ID"
        },
        {
        name: "artist_name",
        display: "Artist Name"
        },
        {
        name: "album_title",
        display: "Album Title"
        },
        {
        name: "track_name",
        display: "Track Name"
        },
        {
        name: "track_composer",
        display: "Track Composer"
        },
        {
        name: "track_milliseconds",
        display: "Track Milliseconds"
        },
        {
        name: "genre_name",
        display: "Genre Name"
        },
        {
        name: "media_type_name",
        display: "Media Type Name"
        },

        {
        name: "playlist_name",
        display: "Playlist Name"
        }
    ]
    // jsx below
    return (
        <div className="table-container">
            {<BootstrapTable tableClass={"table"} headers = {true} columns={artistColumns} data={artistData} />}
        </div>
      );
    }
    
    export default Artist;
  