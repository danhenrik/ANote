import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import PropTypes from "prop-types";
import {
  ContentContainer,
  CommunityCardContainer,
  Title,
} from "./CommunityCard.styled"; // Adjust the import for your community card styles
import Tags from "../../Tags/TagsList";
import { useNavigate } from "react-router-dom";
import { Link } from "@mui/material";
import useCommunities from "../../../../api/useCommunities";

const CommunityCard = ({ community }) => {
  const communitiesApi = useCommunities();

  const followCommunity = () => {
    communitiesApi.followCommunity(community.Id);
  };
  const unfollowCommunity = () => {
    communitiesApi.unfollowCommunity(community.Id);
  };

  const navigationHandler = (event) => {
    if (event.target.closest("a")) return;
    if (event.target.closest("#community-follow")) return;
    navigate(`/community/${community.Id}`);
  };
  const navigate = useNavigate();
  return (
    <CommunityCardContainer
      onClick={navigationHandler}
      sx={{ minWidth: "350px", maxWidth: "350px" }}
    >
      <Link onClick={followCommunity} id='community-follow'>
        follow
      </Link>
      <CardContent>
        <Title variant='h7' component='div'>
          {community.Name}
        </Title>
        <Typography color='textSecondary'>{community.Name}</Typography>
        <ContentContainer>abc</ContentContainer>
        <ContentContainer sx={{ marginTop: "10px" }}>
          {community.Tags && <Tags tags={community.Tags} />}
        </ContentContainer>
      </CardContent>
    </CommunityCardContainer>
  );
};

const communityShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Name: PropTypes.string.isRequired,
  Tags: PropTypes.arrayOf(PropTypes.string),
});

CommunityCard.propTypes = {
  community: communityShape.isRequired,
};

export default CommunityCard;
