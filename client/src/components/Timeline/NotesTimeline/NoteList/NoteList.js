import PropTypes from "prop-types";
import { ButtonBase, Grid } from "@mui/material";
import NoteCard from "../NoteCard/NoteCard";
import { useModal } from "../../../../store/modal-context";
import { useAuth } from "../../../../store/auth-context";
import NoteForm from "../NoteForm/NoteForm";
import LoginForm from "../../../AccessControl/Login/LoginForm";
import TimelineList from "../../TimelineList";
import useCommunities from "../../../../api/useCommunities";
import { useEffect, useState } from "react";

const NoteList = ({ notes, communityId, setNotesHandler }) => {
  const modal = useModal();
  const auth = useAuth();
  const communitiesApi = useCommunities();
  const [isFollowing, setIsFollowing] = useState(false);

  useEffect(() => {
    if (auth.isAuthenticated) {
      const userFollowedCommunities = async () => {
        const communities = await communitiesApi.fetchCommunitiesByUser();
        if (communities) {
          if (communities.some((community) => community.Id === communityId)) {
            setIsFollowing(true);
          }
        }
      };
      userFollowedCommunities();
    }
  }, []);

  const followCommunity = async () => {
    const response = communitiesApi.followCommunity(communityId);
    if (response) {
      setIsFollowing(true);
    }
  };

  const handleAddNoteModal = () => {
    auth.isAuthenticated || isFollowing
      ? modal.openModal(
          auth.isAuthenticated ? (
            <NoteForm
              notes={notes}
              communityId={communityId}
              closeModal={modal.closeModal}
              setNotesHandler={setNotesHandler}
            ></NoteForm>
          ) : (
            <LoginForm closeModal={modal.closeModal}></LoginForm>
          )
        )
      : followCommunity(communityId);
  };
  //trocar botao para um de seguir e um de adicionar, deve aparecer adicionar para usuario deslogado
  const buttonText = !isFollowing ? "Seguir Comunidade" : "Adicionar Nota";

  return (
    <TimelineList
      handleAddModal={handleAddNoteModal}
      addButtonText={buttonText}
    >
      {notes.map((note) => (
        <Grid item key={note.Id}>
          <ButtonBase>
            <NoteCard note={note} />
          </ButtonBase>
        </Grid>
      ))}
    </TimelineList>
  );
};

const noteShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Title: PropTypes.string.isRequired,
  Content: PropTypes.string.isRequired,
  //LikesCount: PropTypes.number.isRequired,
  //Likes: PropTypes.arrayOf(PropTypes.string).isRequired,
  PublishedDate: PropTypes.string.isRequired,
  UpdatedDate: PropTypes.string.isRequired,
  Author: PropTypes.string.isRequired,
  Tags: PropTypes.arrayOf(PropTypes.string),
  Communities: PropTypes.arrayOf(PropTypes.string),
  //CommentCount: PropTypes.number.isRequired,
  //Commenters: PropTypes.arrayOf(PropTypes.string).isRequired,
});

NoteList.propTypes = {
  notes: PropTypes.arrayOf(noteShape).isRequired,
  communityId: PropTypes.any,
  setNotesHandler: PropTypes.func.isRequired,
};

export default NoteList;
