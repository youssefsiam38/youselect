import React, { Fragment, useState } from 'react'
import { NativeSelect, FormControl, Input, InputLabel, FormHelperText } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import Select from '@material-ui/core/Select';

const useStyles = makeStyles((theme) => ({
    formControl: {
        margin: theme.spacing(1),
        minWidth: 120,
    },
    selectEmpty: {
        marginTop: theme.spacing(2),
    },
}));
function Form() {
    const classes = useStyles();
    const [state, setState] = useState({
        store: 'all',
    });

    const handleChange = (event) => {
        const name = event.target.name;
        setState({
            ...state,
            [name]: event.target.value,
        });
    };
    return (
        <Fragment>
            
            <FormControl>
                <InputLabel htmlFor="title">Title</InputLabel>
                <Input id="title" aria-describedby="my-helper-text" />
            </FormControl>
            <FormControl className={classes.formControl}>
                <InputLabel htmlFor="category">Category</InputLabel>
                <Select
                    native
                    value={state.age}
                    onChange={handleChange}
                    inputProps={{
                        name: 'category',
                        id: 'category',
                    }}
                >
                    <option aria-label="None" value="" />
                    <option value={10}>Ten</option>
                    <option value={20}>Twenty</option>
                    <option value={30}>Thirty</option>
                </Select>
            </FormControl>
            {/* <br></br> */}
            <FormControl className={classes.formControl}>
                <InputLabel htmlFor="category">Category</InputLabel>
                <Select
                    native
                    value={state.age}
                    onChange={handleChange}
                    inputProps={{
                        name: 'category',
                        id: 'category',
                    }}
                >
                    <option aria-label="None" value="" />
                    <option value={10}>Ten</option>
                    <option value={20}>Twenty</option>
                    <option value={30}>Thirty</option>
                </Select>
            </FormControl>
        </Fragment>
    )
}

export default Form