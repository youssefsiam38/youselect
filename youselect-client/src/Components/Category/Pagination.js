import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Pagination from '@material-ui/lab/Pagination';
import { Grid } from '@material-ui/core'
import { useHistory } from "react-router-dom";
import updateQueryParameter from '../../Utils/updateQueryParameter.js'


const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
      marginTop: theme.spacing(2),
    }
  },
}));

export default function BasicPagination({ pages }) {
  const classes = useStyles();
  const history = useHistory()

  const onChange = (e, p) => {
    history.push(updateQueryParameter("page", p))
}

  return (
    <div className={classes.root}>
      <Grid container justify="center" alignItems="center">
        <Grid item >
          <Pagination style={{ marginTop: "20px", marginBottom: "20px" }} onChange={onChange} count={pages} />
        </Grid>
      </Grid>
    </div>
  );
}