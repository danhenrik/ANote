import TextField from "@mui/material/TextField";
import { useField } from "formik";
import PropTypes from "prop-types";
import { Autocomplete, createFilterOptions } from "@mui/material";

const OPTIONS_LIMIT = 5;
const defaultFilterOptions = createFilterOptions();

const filterOptions = (options, state) => {
  return defaultFilterOptions(options, state).slice(0, OPTIONS_LIMIT);
};

const InputAutocomplete = ({
  name,
  options,
  addToList,
  setInputValue,
  ...props
}) => {
  const [field, meta] = useField(name);

  const handleInputKeyPress = (event) => {
    if (event.key === "Enter") {
      event.preventDefault();
    }
  };
  return (
    <Autocomplete
      freeSolo
      style={{ width: "100%" }}
      filterOptions={filterOptions}
      {...field}
      options={options}
      isOptionEqualToValue={(option, value) => {
        return option === value;
      }}
      getOptionLabel={(option) => option}
      onChange={(_, value) => {
        field.onChange({
          target: { name, value: value },
        });
        value && addToList(value);
      }}
      renderInput={(params) => (
        <TextField
          onKeyDown={handleInputKeyPress}
          value={null}
          {...params}
          {...props}
          error={Boolean(meta.touched && meta.error)}
          fullWidth
          helperText={meta.touched && meta.error}
          name={name}
          id={name}
          variant='outlined'
          onChange={(event) => {
            setInputValue && setInputValue(event.target.value);
          }}
        />
      )}
    />
  );
};

InputAutocomplete.propTypes = {
  name: PropTypes.string.isRequired,
  addToList: PropTypes.func.isRequired,
  options: PropTypes.arrayOf(PropTypes.string).isRequired,
  value: PropTypes.string,
  setInputValue: PropTypes.func,
};

export default InputAutocomplete;
