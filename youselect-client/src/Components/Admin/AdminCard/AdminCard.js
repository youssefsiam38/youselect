import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import { Grid, Typography } from '@material-ui/core';
import { Link } from 'react-router-dom'


const useStyles = makeStyles({
  root: {
    minWidth: 275,
    minHeight: 250
  }
});

export default function SimpleCard(props) {
  const classes = useStyles();

  return (

    <Grid container justify="center" alignItems="center">
      <Grid item>

        <Button  className={classes.root} component={Link} to={props.to} style={{borderRadius: '82px'}} >

          <Typography variant="h2" component="h2">
            {props.children}
          </Typography>

        </Button>
      </Grid>
    </Grid>

  );
}
