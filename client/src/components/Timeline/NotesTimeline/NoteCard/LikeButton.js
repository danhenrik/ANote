import React, { useEffect, useState } from "react";
import FavoriteIcon from "@mui/icons-material/Favorite";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
import axios from "axios";

const LikeButton = (note) => {
  const [likes, setLikes] = useState(0);
  const [isClicked, setIsClicked] = useState(false);
  const [isHovered, setIsHovered] = useState(false);

  const handleClick = async () => {
    try {
      if (isClicked) {
        const deletedLike = await axios.delete(
          "/likes/" + note.note.Author + "/" + note.note.Id
        );
        if (deletedLike) {
          setLikes(likes - 1);
        }
      } else {
        const likeData = {
          user_id: note.note.Author,
          note_id: note.note.Id,
        };
        const postedLike = axios.post("/likes", likeData);
        if (postedLike) {
          setLikes(likes + 1);
        }
      }
      setIsClicked(!isClicked);
    } catch (error) {
      console.log("Like failed:", error);
    }
  };

  useEffect(() => {
    const initLikes = async () => {
      try {
        const response = await axios.get(
          "/likes/" + note.note.Author + "/" + note.note.Id
        );

        const numberLikes = await axios.get("/likes/count/" + note.note.Id);

        if (response.data.data) setIsClicked(true);
        setLikes(numberLikes.data.data);
      } catch (error) {
        console.log("Likes retrieving failed: ", error);
      }
    };

    initLikes();
  }, []);

  const favoriteIconStyling = {
    color: "red",
    ...(isHovered && { transform: "scale(1.2)" }),
  };

  return (
    <>
      <span
        onClick={handleClick}
        onMouseEnter={() => setIsHovered(true)}
        onMouseLeave={() => setIsHovered(false)}
      >
        {isClicked || isHovered ? (
          <FavoriteIcon
            id='like-button'
            className={`like-button ${isClicked && "liked"}`}
            style={favoriteIconStyling}
          />
        ) : (
          <FavoriteBorderIcon
            id='like-button'
            className={`like-button ${isClicked && "liked"}`}
            style={{ marginRight: "2px", color: "red" }}
          />
        )}
      </span>
      <span
        style={{
          color: "red",
          marginTop: "4px",
        }}
      >
        {likes}
      </span>
    </>
  );
};

export default LikeButton;
