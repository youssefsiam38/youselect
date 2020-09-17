import React from 'react';
import AdminCard from './AdminCard/AdminCard.js'
import ProductsTable from './ProductsTable/ProductsTable.js'
import CategoriesTable from './CategoriesTable/CategoriesTable.js'
import { Grid } from '@material-ui/core';


const Admin = function () {
    return (
        <React.Fragment>
            <br/>
            <Grid container alignItems="center" justify="space-evenly">

                <Grid item xs md={5}>

                    <AdminCard to="/admin/products">
                        Products
                    </AdminCard>
                </Grid>
                <Grid item xs md={5}>
                    <AdminCard to="/admin/categories">
                        Categories
                    </AdminCard>
                </Grid>
                <Grid item xs md={5}>
                    <AdminCard to="/admin/stores">
                        Stores
                    </AdminCard>
                </Grid>
            </Grid>
        </React.Fragment>
    )
}

export default Admin