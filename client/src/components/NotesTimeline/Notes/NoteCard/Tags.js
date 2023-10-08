import styled from "@mui/system/styled";
import { Link } from "react-router-dom";
import PropTypes from "prop-types";

const TagList = styled("div")`
  display: flex;
  flex-wrap: wrap;
`;

const TagLink = styled(Link)`
  text-decoration: none;
  background-color: #7f56d9;
  color: #fff;
  border: none;
  border-radius: 4px;
  margin: 4px;
  padding: 8px 18px;
  cursor: pointer;

  &:hover {
    text-decoration: underline; /* Apply underline on hover */
  }
`;

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
