import React, { Fragment } from 'react'
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
    const [state, setState] = React.useState({
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
                <InputLabel htmlFor="min">Minimum Price($)</InputLabel>
                <Input type="number" inputProps={{ min: '0' }} id="min" aria-describedby="my-helper-text" />
                <FormHelperText id="my-helper-text">We'll never share your email.</FormHelperText>
            </FormControl>
            <FormControl>
                <InputLabel htmlFor="max">Maximum Price($)</InputLabel>
                <Input type="number" inputProps={{ min: '1' }} id="max" aria-describedby="my-helper-text" />
                <FormHelperText id="my-helper-text">We'll never share your email.</FormHelperText>
            </FormControl>
            <br></br>
            <FormControl className={classes.formControl}>
                <InputLabel htmlFor="store">Only on (store)</InputLabel>
                <NativeSelect
                    id="store"
                    value={state.store}
                    onChange={handleChange}
                    name="store"
                    className={classes.selectEmpty}
                    inputProps={{ 'aria-label': 'age' }}
                >
                    <option value="all">All</option>
                    <option value="amazon">Amazon</option>
                    <option value="ebay">ebay</option>
                    <option value="alibaba">alibaba</option>
                </NativeSelect>
                <FormHelperText>With visually hidden label</FormHelperText>
            </FormControl>
        </Fragment>
    )
}

export default Form