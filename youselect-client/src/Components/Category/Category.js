
import React, { Fragment, useState, useEffect } from 'react'
import Filter from './Filter/Filter.js'
import { Grid, Typography } from '@material-ui/core'
import Card from './Card.js'
import { useParams } from 'react-router-dom'
import Pagination from './Pagination.js'



function Category(props) {
    const { categ } = useParams()
    const [products, setProducts] = useState([
        { id: 1, title: '32-Inch HD LED Standard TV ATA 32 Black', category: 'usbs', price: 230, store: 'noon', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61Ck0CO5%2BHL._AC_SL1500_.jpg', priority: 3 },
        { id: 4, title: '32-Inch HD LED Standard TV ATA 32 Black', category: 'usbs', price: 230, store: 'noon', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61Ck0CO5%2BHL._AC_SL1500_.jpg', priority: 3 },
        { id: 5, title: '32-Inch HD LED Standard TV ATA 32 Black', category: 'usbs', price: 230, store: 'noon', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61Ck0CO5%2BHL._AC_SL1500_.jpg', priority: 3 },
        { id: 6, title: '32-Inch HD LED Standard TV ATA 32 Black', category: 'usbs', price: 230, store: 'noon', imageURL: 'https://images-na.ssl-images-amazon.com/images/I/61Ck0CO5%2BHL._AC_SL1500_.jpg', priority: 3 },
        { id: 2, title: '32-Inch HD LED Standard TV ATA 32 white', category: 'usbs', price: 230, store: 'noon', imageURL: 'https://k.nooncdn.com/t_desktop-pdp-v1/v1583746577/N35521717A_1.jpg', affiliateURL: 'https://www.noon.com/egypt-en/32-inch-hd-led-standard-tv-ata-32-black/N23204693A/p?o=e725c564befce2ee', productURL: 'https://www.noon.com/egypt-en/32-inch-hd-led-standard-tv-ata-32-black/N23204693A/p?o=e725c564befce2ee', priority: 3 },
        { id: 3, title: '32-Inch HD LED Standard TV ATA 32 Black', category: 'usbs', price: 230, store: 'noon', imageURL: 'https://imgaz3.staticbg.com/thumb/large/oaupload/banggood/images/5C/BA/bf016515-91e1-4bc6-8d96-54a5062386ff.jpg.webp', affiliateURL: 'https://www.noon.com/egypt-en/32-inch-hd-led-standard-tv-ata-32-black/N23204693A/p?o=e725c564befce2ee', productURL: 'https://www.noon.com/egypt-en/32-inch-hd-led-standard-tv-ata-32-black/N23204693A/p?o=e725c564befce2ee', priority: 3 }
    ])
    const [storeURLs, setStoreURLs] = useState({
        noon: {
            affURL: 'https://www.noon.com/egypt-en/',
            imageURL: 'https://li0.rightinthebox.com/images/dfp/fs-images/2020/be41d3d012451e7fdde71cbcc34e6972.jpg'
        }
    })


    const getStoreURL = (storeName) => {

    }

    const smallDisplay = window.screen.width < 950
    return (
        <Fragment>
            <br></br>
            <br></br>
            <br></br>
            <br></br>
            <br></br>
            <Grid container alignItems="center" justify="space-between">
                {smallDisplay ? 
                    (<>
                        <Grid item xs={12}  zeroMinWidth>
                            <Typography noWrap variant="h4" style={{ textAlign: 'center' }} >{categ}</Typography>
                        </Grid>
                        <Grid item xs={12}>
                            <Filter>filter</Filter>
                        </Grid>
                    </>) : 
                    (<>
                        <Grid item xs={3}>
                            <Filter>filter</Filter>
                        </Grid>
                        <Grid item xs={6} zeroMinWidth>
                            <Typography noWrap variant="h4" style={{ textAlign: 'center' }} >{categ}</Typography>
                        </Grid>
                    </>)
                }
                <Grid item xs={12} md={3}></Grid>
            </Grid>

            <br></br>
            <br></br>
            <Grid container alignItems="center" justify="space-evenly" spacing={3}>
                {products.map((product) => (
                    <Card
                        key={product.id}
                        product={product}
                        storeAffURL={storeURLs[product.store].affURL}
                        storeImageURL={storeURLs[product.store].imageURL}
                    />
                ))

                }
            </Grid>
            <Pagination pages={15}/>
        </Fragment>

    )
}

export default Category