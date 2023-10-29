import CardContent from "@mui/material/CardContent";
import PropTypes from "prop-types";
import { CommunityCardContainer, Title } from "./CommunityCard.styled"; // Adjust the import for your community card styles
import { useNavigate } from "react-router-dom";
import { Link } from "@mui/material";
import useCommunities from "../../../../api/useCommunities";
import { useState } from "react";

const CommunityCard = ({ community, isFollowing, communityFollowHandler }) => {
  const communitiesApi = useCommunities();
  const [isFollowed, setIsFollowed] = useState(isFollowing);

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
    navigate(`/communities/${community.Id}`);
  };
  const navigate = useNavigate();
  return (
    <CommunityCardContainer
      onClick={navigationHandler}
      sx={{ minWidth: "350px", maxWidth: "350px", margin: "10px" }}
    >
      {!isFollowed ? (
        <Link onClick={followCommunity} id='community-follow'>
          follow
        </Link>
      ) : (
        <Link onClick={unfollowCommunity} id='community-follow'>
          unfollow
        </Link>
      )}
      <CardContent>
        <Title style={{ textAlign: "center" }} variant='h7' component='div'>
          {community.Name}
        </Title>
      </CardContent>
    </CommunityCardContainer>
  );
};

const communityShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Name: PropTypes.string.isRequired,
});

CommunityCard.propTypes = {
  community: communityShape.isRequired,
  isFollowing: PropTypes.bool.isRequired,
  communityFollowHandler: PropTypes.func,
};

export default CommunityCard;
