import PropTypes from "prop-types";
import { TagBlock, TagLink, TagList } from "./TagsList.styled";
import { IconButton } from "@mui/material";
import CancelOutlinedIcon from "@mui/icons-material/CancelOutlined";

const Tags = ({ tags, hasLink, hasDelete, deletionHandler }) => {
  return (
    <TagList>
      {tags.map((tag, index) =>
        !hasLink ? (
          <TagBlock sx={{ height: "50px" }} key={index}>
            {tag}
            {hasDelete && (
              <IconButton
                sx={{ marginLeft: "3px" }}
                onClick={() => deletionHandler(tag)}
              >
                <CancelOutlinedIcon
                  sx={{
                    fontSize: "1.4rem",
                    color: "white",
                    "&:hover": {
                      color: "red",
                    },
                  }}
                />
              </IconButton>
            )}
          </TagBlock>
        ) : (
          <TagLink key={index} to={`/tags/${tag}`}>
            {tag}
          </TagLink>
        )
      )}
    </TagList>
  );
};

Tags.propTypes = {
  hasLink: PropTypes.bool,
  hasDelete: PropTypes.bool,
  deletionHandler: PropTypes.func,
  tags: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default Tags;
