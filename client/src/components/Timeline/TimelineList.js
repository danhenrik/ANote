import { Grid } from "@mui/material";
import { CreateButton, GridWrapper } from "./TimelineList.styled";
import PropTypes from "prop-types";

const TimelineList = ({ children, handleAddModal, addButtonText }) => {
  return (
    <>
      <GridWrapper>
        <Grid
          container
          spacing={{ xs: 2, md: 3 }}
          columns={{ xs: 1, sm: 1, md: 3 }}
        >
          <Grid
            item
            key='button'
            sx={{ display: "flex", justifyContent: "flex-end", width: "100%" }}
          >
            <CreateButton onClick={handleAddModal}>
              {addButtonText}
            </CreateButton>
          </Grid>
          {children}
        </Grid>
      </GridWrapper>
    </>
  );
};

TimelineList.propTypes = {
  children: PropTypes.node.isRequired,
  handleAddModal: PropTypes.func.isRequired,
  addButtonText: PropTypes.string.isRequired,
};

export default TimelineList;
