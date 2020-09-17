import React, {useEffect} from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import CssBaseline from "@material-ui/core/CssBaseline";
import Box from "@material-ui/core/Box";
import Container from "@material-ui/core/Container";
import Fab from "@material-ui/core/Fab";
import KeyboardArrowUpIcon from "@material-ui/icons/KeyboardArrowUp";
import { fade, makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete';
import ScrollTop from './ScrollTop.js';
import Grid from '@material-ui/core/Grid';
import SearchSharpIcon from '@material-ui/icons/SearchSharp';
import IconButton from '@material-ui/core/IconButton';
import { Link } from 'react-router-dom'
import updateQueryParameter from '../../Utils/updateQueryParameter.js'
import { useHistory } from "react-router-dom";



console.log(window.location.href.split('/'))

export default function Nav(props) {
    const history = useHistory()
    const useStyles = makeStyles((theme) => ({
        title: {
            flexGrow: 1,
            display: 'none',
            [theme.breakpoints.up('sm')]: {
                display: 'block',
            },
        },
        search: {
            position: 'relative',
            borderRadius: theme.shape.borderRadius,
            backgroundColor: fade(theme.palette.common.white, 0.15),
            '&:hover': {
                backgroundColor: fade(theme.palette.common.white, 0.25),
            },
            marginLeft: 0,
            width: '100%',
            [theme.breakpoints.up('sm')]: {
                marginLeft: theme.spacing(1),
                width: 'auto',
            },
        }
    }));

    useEffect(() => {
        // document.querySelector('#search').focus()
    })


    const FancyLink = React.forwardRef((props, ref) => (
        <a href={ref}> <Button variant="outlined" color="inherit">{props.children}</Button></a>
    ))
    const smallDisplay = window.screen.width < 422
    const classes = useStyles()
    return (
        <React.Fragment>
            <CssBaseline />
            <AppBar style={{ paddingBottom: '6px', backgroundColor: 'green' }} >

                <Toolbar>
                    <Grid container alignItems="center" justify="space-between">
                        <Grid item md={7} xs={12}>
                            <Button component={Link} to="/" color="inherit" style={{ borderRadius: '15px' }} >
                                <img height="50" width="50" style={{ height: '100%' }} src="/l.png" />
                            </Button>
                            {!smallDisplay ?
                                (<Button style={{ marginRight: "15px", marginLeft: "15px", fontSize: "0.8rem" }} component={Link} to="/about" variant="outlined" color="inherit">Shop</Button>) : null
                            }
                            <Button style={{ marginRight: "15px", marginLeft: smallDisplay ? "15px" : "0", fontSize: "0.8rem" }} component={Link} to="/about" variant="outlined" color="inherit">Categories</Button>
                            {window.screen.width > 315 ?
                                (<Button style={{ fontSize: "0.8rem" }} component={Link} to="/about" variant="outlined" color="inherit">About</Button>) : null
                            }
                        </Grid>
                        <Grid item md={5} xs={12}>
                            <form onSubmit={(e) => {
                                const value = document.querySelector("#search").value
                                e.preventDefault()
                                history.push(updateQueryParameter("s", value))
                                history.push(updateQueryParameter("page", 1))
                            }}>
                                <Grid container>
                                    <Grid item sm={11} xs={10}>
                                        <div className={classes.search}>
                                            <Autocomplete                                                
                                                freeSolo
                                                id="search"
                                                value="hey"
                                                options={["1", "2"]}
                                                renderInput={(params) => <TextField {...params} label="Search" variant="outlined" />}
                                                style={{ backgroundColor: 'white', borderRadius: '10px' }}
                                            />
                                        </div>
                                    </Grid>
                                    <Grid item xs={1}>
                                        <Button type="submit" variant="text" color="inherit" fullWidth style={{ textAlign: 'center', height: '100%', borderRadius: '1000px' }}>
                                            <SearchSharpIcon fontSize="large" style={{ width: "100%", heigth: "100%" }} />
                                        </Button>
                                    </Grid>
                                </Grid>
                            </form>
                        </Grid>
                    </Grid>
                </Toolbar>
            </AppBar>
            <Toolbar id="back-to-top-anchor" />
            <Container>
                <Box my={2}>
                    {props.children}
                </Box>
            </Container>
            <ScrollTop {...props}>
                <Fab style={{ backgroundColor: 'green', color: 'white' }} size="small" aria-label="scroll back to top">
                    <KeyboardArrowUpIcon />
                </Fab>
            </ScrollTop>
        </React.Fragment>
    );

}
