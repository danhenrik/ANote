import CardContent from "@mui/material/CardContent";
import PropTypes from "prop-types";
import {
  CommunityCardContainer,
  FollowButton,
  FollowButtonFollowing,
  Title,
} from "./CommunityCard.styled"; // Adjust the import for your community card styles
import { useNavigate } from "react-router-dom";
import useCommunities from "../../../../api/useCommunities";
import { useState } from "react";
import { useAuth } from "../../../../store/auth-context";
import randomColor from "randomcolor";
import { AvatarBackground } from "../../NotesTimeline/NoteCard/NoteCard.styled";

const CommunityCard = ({ community, isFollowing, communityFollowHandler }) => {
  const [randomColorElement] = useState(
    randomColor({ luminosity: "light", format: "rgb" })
  );
  const communitiesApi = useCommunities();
  const [isFollowed, setIsFollowed] = useState(isFollowing);
  const auth = useAuth();

  const followCommunity = () => {
    if (communitiesApi.followCommunity(community.Id)) {
      setIsFollowed(true);
      communityFollowHandler(community.Id);
    }
  };
  const unfollowCommunity = () => {
    if (communitiesApi.unfollowCommunity(community.Id)) {
      setIsFollowed(false);
      communityFollowHandler(community.Id);
    }
  };

  const navigationHandler = (event) => {
    if (event.target.closest("a")) return;
    if (event.target.closest("#community-follow")) return;
    navigate(`/communities/${community.Name}`);
  };
  const navigate = useNavigate();
  return (
    <CommunityCardContainer
      onClick={navigationHandler}
      sx={{
        minWidth: "350px",
        maxWidth: "350px",
        margin: "10px",
      }}
    >
      <div
        style={{
          minWidth: "350px",
          maxWidth: "350px",
          display: "flex",
        }}
      >
        {auth.isAuthenticated &&
          (!isFollowed ? (
            <FollowButton
              onClick={(e) => {
                e.preventDefault(); // Prevent the default action
                followCommunity();
              }}
              id='community-follow'
            >
              Seguir
            </FollowButton>
          ) : (
            <FollowButtonFollowing
              onClick={(e) => {
                e.preventDefault(); // Prevent the default action
                unfollowCommunity();
              }}
              id='community-follow'
            >
              Seguindo
            </FollowButtonFollowing>
          ))}
      </div>

      <CardContent sx={{ padding: "5px" }}>
        <AvatarBackground
          style={{
            height: "40px",
            display: "flex",
            alignItems: "center", // Center vertically
            justifyContent: "center", // Center horizontally
          }}
          randomColor={randomColorElement}
        >
          <Title
            style={{ textTransform: "uppercase" }}
            variant='h5'
            component='div'
          >
            {community.Name}
          </Title>
        </AvatarBackground>

        <img
          src={community.Background}
          alt={community.Name}
          style={{
            top: 0,
            left: 0,
            width: "100%", // Adjust the width and height as needed
            height: "100%",
          }}
        />
      </CardContent>
    </CommunityCardContainer>
  );
};

const communityShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Name: PropTypes.string.isRequired,
  Background: PropTypes.any.isRequired,
});

CommunityCard.propTypes = {
  community: communityShape.isRequired,
  isFollowing: PropTypes.bool.isRequired,
  communityFollowHandler: PropTypes.func,
};

export default CommunityCard;
