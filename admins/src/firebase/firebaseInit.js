import firebase from "firebase/app"
import "firebase/firestore"


const firebaseConfig = {
  apiKey: "AIzaSyBOdHQX6J6yofaN2-EmR2bWPezyTUqDRLw",
  authDomain: "blogs-dfc4e.firebaseapp.com",
  projectId: "blogs-dfc4e",
  storageBucket: "blogs-dfc4e.appspot.com",
  messagingSenderId: "1073460183326",
  appId: "1:1073460183326:web:40d49c6dd0df024c65890f",
  measurementId: "G-KR1SM6T51W",
};


const firebaseApp = firebase.initializeApp(firebaseConfig);

const timestamp = firebase.firestore.FieldValue.severTimestamp;


export { timestamp };

export default firebaseApp.firestore();