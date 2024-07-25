This micro-service is part of course-registration system.

This service manages offered courses for professors and registered courses for students using MongoDB
Every successful request would return Status: 200 with requested data.
In case of any error encoutered, {response: "error message"} would be returned as request's response with respective status code.

<h1>API Endpotins</h1>

<h3>Offered course</h3>
<hr>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>POST</td>
    <td>/offered_course</td>
    <td>Offers new course <br> Fields: <br> <i>Course_id<br>CRN<br>Offered_by<br>Day_Time consisting of day, start_time and end_time </i></td>
  </tr>
  <tr>
    <td>GET</td>
    <td>/offered_course <br>/offered_course?course_id=? <br> /offered_course?email_id=? <br> /offered_course?crn=?</td>
    <td>Fetches all offered courses or based on offered_course_id, offerring_professor_email_id or offered_course's CRN respectively</td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/offered_course/:crn</td>
    <td>Update offered_course using the given CRN <br>Fields:<br><i>Day_time consisting of day, start_time and end_time</i></td>
  </tr>
  <tr>
    <td>DELETE</td>
    <td>/offered_course/:crn</td>
    <td>Delete an offered course using the given CRN course</td>
  </tr>
</table>
