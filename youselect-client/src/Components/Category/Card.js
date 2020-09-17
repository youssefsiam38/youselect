import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';


const useStyles = makeStyles({
  root: {
    maxWidth: 345,
    margin: 'auto'
  },
  media: {
    height: 350,
    maxWidth: 300,
  },
});

export default function MediaCard({ product, storeAffURL, storeImageURL }) {
  const classes = useStyles();

  return (
    <Card className={classes.root} style={{ margin: '10px' }}>
      <CardActionArea onClick={() => window.open(product.affiliateURL)}>
        <CardMedia>
          <Grid container justify="center" alignItems="stretch">

            <Grid item>

              <img
                src={product.imageURL}
                // width="300"
                className={classes.media}
                title={product.title}
              />
            </Grid>
          </Grid>
        </CardMedia>
        <CardContent>
          <Typography variant="body2" color="textPrimary" component="p">
            {product.title}
          </Typography>
        </CardContent>
      </CardActionArea>
      <CardActions>
        <Grid container justify="space-between" alignItems="center">
          <Grid item xs>
            <Typography gutterBottom variant="body1" component="h4">
              Price: {product.price} <span style={{ fontSize: 'smaller' }}>USD</span>
            </Typography>
          </Grid>
          <Grid item xs>
            <a href={storeAffURL} >
              <img style={{ height: '55px', width: "140px" }} src={storeImageURL} ></img>
            </a>
          </Grid>
        </Grid>
      </CardActions>
    </Card>
  );
}