import PropTypes from "prop-types";
import { ButtonBase, Grid } from "@mui/material";
import CommunityForm from "../CommunityForm/CommunityForm";
import { useModal } from "../../../../store/modal-context";
import { useAuth } from "../../../../store/auth-context";
import TimelineList from "../../TimelineList";
import LoginForm from "../../../AccessControl/Login/LoginForm";
import CommunityCard from "../CommunityCard/CommunityCard";

const CommunityList = ({ communities }) => {
  const modal = useModal();
  const auth = useAuth();

  const handleAddCommunityModal = () => {
    modal.openModal(
      auth.isAuthenticated ? (
        <CommunityForm
          communities={communities}
          closeModal={modal.closeModal}
        />
      ) : (
        <LoginForm closeModal={modal.closeModal}></LoginForm>
      )
    );
  };
  const buttonText = "Adicionar Comunidade";

  return (
    <TimelineList
      handleAddModal={handleAddCommunityModal}
      addButtonText={buttonText}
    >
      {communities.map((community) => (
        <Grid item key={community.Id}>
          <ButtonBase>
            <CommunityCard community={community}></CommunityCard>
          </ButtonBase>
        </Grid>
      ))}
    </TimelineList>
  );
};

const communityShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Name: PropTypes.string.isRequired,
});

CommunityList.propTypes = {
  communities: PropTypes.arrayOf(communityShape).isRequired,
};

export default CommunityList;
