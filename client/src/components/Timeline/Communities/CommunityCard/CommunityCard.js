import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import PropTypes from "prop-types";
import {
  ContentContainer,
  CommunityCardContainer,
  StyledLink,
  Title,
} from "./CommunityCard.styled"; // Adjust the import for your community card styles
import Tags from "../../Tags/TagsList";

const CommunityCard = ({ community }) => {
  return (
    <CommunityCardContainer sx={{ minWidth: "350px", maxWidth: "350px" }}>
      <CardContent>
        <Title variant='h7' component='div'>
          {community.Name}
        </Title>
        <StyledLink to='/community/{community.Id}'>
          <Typography color='textSecondary'>{community.Name}</Typography>
        </StyledLink>
        <ContentContainer>abc</ContentContainer>
        <ContentContainer sx={{ marginTop: "10px" }}>
          <Tags tags={community.Tags} />
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
