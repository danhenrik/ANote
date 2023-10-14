import PropTypes from "prop-types";
import { TagLink, TagList } from "./Tags.styled";

function Tags({ tags }) {
  return (
    <TagList>
      {tags.map((tag, index) => (
        <TagLink key={index} to={`/tags/${tag}`}>
          {tag}
        </TagLink>
      ))}
    </TagList>
  );
}

Tags.propTypes = {
  tags: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default Tags;
