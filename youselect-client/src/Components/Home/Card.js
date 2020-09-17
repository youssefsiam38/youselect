import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { Link } from 'react-router-dom'
import Grid from '@material-ui/core/Grid';



const useStyles = makeStyles({
  root: {
    maxWidth: 345,
    margin: 'auto'
  },
  media: {
    height: 200,
    maxWidth: 300,
  },
});

export default function MediaCard(props) {
  const classes = useStyles();

  return (
    <Card className={classes.root} style={{margin: '10px'}}>
      <CardActionArea component={Link} to="/laptop">
      <CardMedia>
          <Grid container justify="center" alignItems="stretch">

            <Grid item>

              <img
                src={props.imageURL}
                // width="300"
                className={classes.media}
                title={props.title}
              />
            </Grid>
          </Grid>
        </CardMedia>
        <CardContent>
          <Typography align="center" gutterBottom variant="h5" component="h2">
            {props.title}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}