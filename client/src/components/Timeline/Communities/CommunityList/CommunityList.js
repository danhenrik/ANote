import PropTypes from "prop-types";
import { ButtonBase, Grid } from "@mui/material";
import CommunityForm from "../CommunityForm/CommunityForm";
import { useModal } from "../../../../store/modal-context";
import { useAuth } from "../../../../store/auth-context";
import TimelineList from "../../TimelineList";
import LoginForm from "../../../AccessControl/Login/LoginForm";
import CommunityCard from "../CommunityCard/CommunityCard";
import useCommunities from "../../../../api/useCommunities";
import { useEffect, useState } from "react";
import EmptyNotes from "../../NotesTimeline/EmptyNotes";

const CommunityList = ({ communities, setCommunitiesHandler }) => {
  const [userCommunities, setUserCommunities] = useState([]);
  const modal = useModal();
  const auth = useAuth();
  const communitiesApi = useCommunities();

  const communityFollowHandler = (id) => {
    setCommunitiesHandler((communities) =>
      communities.filter((community) => community.Id !== id)
    );
  };

  useEffect(() => {
    const fetchUserCommunities = async () => {
      try {
        const userFollowedCommunities =
          await communitiesApi.fetchCommunitiesByUser();
        setUserCommunities(userFollowedCommunities);
      } catch (error) {
        console.error("Error fetching user's followed communities:", error);
      }
    };
    if (auth.isAuthenticated) {
      fetchUserCommunities();
    } else {
      setUserCommunities([]);
    }
  }, [communities, auth.isAuthenticated]);

  const checkFollowStatus = (communityId) => {
    return userCommunities.some((community) => community.Id === communityId);
  };

  const handleAddCommunityModal = () => {
    modal.openModal(
      auth.isAuthenticated ? (
        <CommunityForm
          communities={communities}
          closeModal={modal.closeModal}
          setCommunitiesHandler={setCommunitiesHandler}
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
      {communities && communities.length ? (
        communities.map((community) => (
          <Grid item key={community.Id}>
            <ButtonBase>
              <CommunityCard
                isFollowing={checkFollowStatus(community.Id)}
                community={community}
                communityFollowHandler={communityFollowHandler}
              ></CommunityCard>
            </ButtonBase>
          </Grid>
        ))
      ) : (
        <EmptyNotes clickHandler={handleAddCommunityModal}>
          Nenhuma Comunidade Aqui, Mas Você Pode Criar ou Seguir uma
        </EmptyNotes>
      )}
    </TimelineList>
  );
};

const communityShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Name: PropTypes.string.isRequired,
  Background: PropTypes.any.isRequired,
});

CommunityList.propTypes = {
  communities: PropTypes.arrayOf(communityShape).isRequired,
  setCommunitiesHandler: PropTypes.func.isRequired,
};

export default CommunityList;
