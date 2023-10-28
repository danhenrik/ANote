import React from "react";
import TextField from "@mui/material/TextField";
import { useField } from "formik";
import PropTypes from "prop-types";
import { Autocomplete, createFilterOptions } from "@mui/material";

const OPTIONS_LIMIT = 5;
const defaultFilterOptions = createFilterOptions();

const filterOptions = (options, state) => {
  return defaultFilterOptions(options, state).slice(0, OPTIONS_LIMIT);
};

const InputAutocomplete = ({ name, options, fieldName, ...props }) => {
  const [field, meta] = useField(name);
  return (
    <Autocomplete
      filterOptions={filterOptions}
      {...field}
      options={options}
      isOptionEqualToValue={(option, value) => option.id === value.id}
      getOptionLabel={(option) =>
        typeof option === "string" ? option : option[fieldName]
      }
      onChange={(_, value) => {
        field.onChange({
          target: { name, value: value[fieldName] || "" },
        });
      }}
      renderInput={(params) => (
        <TextField
          {...params}
          {...props}
          error={Boolean(meta.touched && meta.error)}
          fullWidth
          helperText={meta.touched && meta.error}
          name={name}
          variant='outlined'
        />
      )}
    />
  );
};

InputAutocomplete.propTypes = {
  name: PropTypes.string.isRequired,
  fieldName: PropTypes.string.isRequired,
  options: PropTypes.arrayOf(
    PropTypes.shape({
      tag: PropTypes.string.isRequired, // Autocomplete option tag
    })
  ).isRequired,
};

export default InputAutocomplete;
