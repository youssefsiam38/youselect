
import React, { useState, useEffect } from 'react'
import { Grid, Typography } from '@material-ui/core'
import Card from './Card.js'
function Home(props) {
    const [categories, setCategories] = useState([
        { id: 1, title: 'anything again', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61FSnPYsXzL._AC_SL1200_.jpg' },
        { id: 2, title: 'anything again', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61FSnPYsXzL._AC_SL1200_.jpg' },
        { id: 3, title: 'anything again', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61FSnPYsXzL._AC_SL1200_.jpg' },
        { id: 4, title: 'anything again', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61FSnPYsXzL._AC_SL1200_.jpg' },
        { id: 5, title: 'anything again', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61FSnPYsXzL._AC_SL1200_.jpg' },
        { id: 6, title: 'anything again', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61FSnPYsXzL._AC_SL1200_.jpg' },
        { id: 7, title: 'anything again', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61FSnPYsXzL._AC_SL1200_.jpg' }
    ])

    return (
        <div>
            <br></br>
            <br></br>
            <br></br>
            <br></br>
            <br></br>
            <Typography noWrap variant="h4" style={{ textAlign: 'center' }} >Pick a Category</Typography>
            <br></br>
            <Grid container alignItems="stretch" justify="flex-start" spacing={3}>
                {categories.map(category => (
                    <Grid key={category.id} item md={4} sm={6} xs={12}>
                        <Card imageURL={category.imageURL} title={category.title}  ></Card>
                    </Grid>
                ))
                }

            </Grid>
        </div>

    )
}

export default Home