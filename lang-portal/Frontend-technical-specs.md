# Frontend Technical Specs 

## Pages 

### Dashboard `/dashboard`

#### Purpose

The purpose of this page is to provide a summary of learning and act as the defautl page when a user visits web app

#### Components 

This page contains the following components 
- Last Study Session
    Shows last activity used 
    shows when last activity used
    summarizes wrong vs correct from last activity has a link to the group 

- Study Progress
    - total words study eg. 3/124
        - across all study session show the total words studied out of all possible words in our database
    - display a mastery progress eg. 0%

- Quick Stats
    - success rate eg. 90%
    - total study sessions eg. 5
    - total active groups eg. 3
    - study streak eg. 3 days

- Start Studying Button 
    - goes to study activites page

#### Needed API Endpoints

we'll need following API endpoints to power this 

- GET api/dashboard/last_study_session
- GET api/dashboard/study_progress
- GET api/dashboard/quick_stats

### Study Activities Index `/study-activities`

#### Purpose

The purpose of this page is to show a collection of study activities with a thumbnail and its name, to either launch or view the study activity.

#### Components 

- Study Activity Card
    - shows a thumbnail of the study activity
    - the name of the study activity
    - a launch button  to take us to the launch page
    - the view page to view more information about past study sessions for this study activity

#### Needed API Endpoints

- GET api/study-activities

### Study Activit Show `/study-activities/:id`

#### Purpose

The purpose of this page is to show a details of  study activity, including past study sessions.

#### Components 
- Name of the study activity
- Thumbnail of the study activity
- Description of the study activity
- Launch botton
- Study Activities Paginated List
    - id
    - activity name 
    - group name 
    - start time
    - end time (inferred by the last word_review_item submitted)
    - number of reviw items

#### Needed API Endpoints

- GET api/study-activities/:id
- GET api/study-activities/:id/study-sessions

### Study Activity Launch `/study-activities/:id/launch`

#### Purpose

The purpose of this page is to launch a study activity.

#### Components
- Name of the study activity
- Launch form
    -  select field for group
    -  launch now button

## Behaviour
After the form is submitted a new tab opens with the study activity based on its URL provided in the database.

Also the after form is submitted the page will redirect to the study session show page

#### Needed API Endpoints

- POST api/study-activities/:id

### Words Index `/words`

#### Purpose

The purpose of this page is to show a collection of words with their romaji, meaning and a part of speech.

#### Components
-paginated word List
    - Columns
        - Japanese
        - Romaji
        - English
        - Correct count
        - Word Count
    - Pagination with 100 items per page
    - Clicking the Japanese word will take us to the word show page 

#### Needed API Endpoints

- GET api/words

### Word Show `/words/:id`

#### Purpose

The purpose of this page is to show information about a specific word.

#### Components
- Japanese
- Romaji
- English
- Study Statistics
    - Correct Count
    - Wrong Count 
- Word Groups
    - show an a series of pills eg. tags
    - when group name is clicked it will take us to the group show page

#### Needed API Endpoints

- GET api/words/:id

### Word Groups Index `/groups`

#### Purpose

The purpose of this page is to show a list of groups in our database.

#### Components
- Paginated Group List
    - Columns
        - Group Name
        - Word Count
    - Clicking the group name will take us to the group show page


#### Needed API Endpoints

- GET api/groups

### Group Show `/groups/:id`

#### Purpose
The purpose of this page is to show information about a specific group.

#### Components
- Group Name
- Group Statistics
    - Total Word Count
- Words in Group (Paginateds List of Words)
    - Should use the same component as the words index page
- Study Sessions (Paginated List of Study Sessions)
    - Should use the same component as the study sessions index page

#### Needed API Endpoints

- GET api/groups/:id (the name and groups stats)
- GET api/groups/:id/words
- GET api/groups/:id/study-sessions

## Study Session Index `/study-sessions`

#### Purpose

The purpose of this page is to show a list of study sessions in our database.

#### Components
- Paginated Study Session List
    - Columns
        - id 
        - Activity Name
        - Group Name
        - Start Time
        - End Time
        - Number of Review Items
    - Clicking the study session id will take us to the study session show page

#### Needed API Endpoints

- GET api/study-sessions


### Study Session Show `/study-sessions/:id`

#### Purpose

The purpose of this page is to show information about a specific study session.

#### Components
- Study Session Information
    - Activity Name
    - Group Name
    - Start Time
    - End Time
    - Number of Review Items
- Words Review Items (Paginated List of Words)
    - Should use the same component as the words index page

#### Needed API Endpoints

- GET api/study-sessions/:id
- GET api/study-sessions/:id/words

### Settings Page `/settings`

#### Purpose

The purpose of this page is to make configuration to the study portal


#### Components
- Theme Selection eg. light, Dark, System Default
- Reset History Button
    - this will delete all study session and word review items
- Full Reset Button
    - this will delete all study session and word review items
    - this will drop all tables and re-create with seed data


#### Needed API Endpoints
- POST api/reset-history
- POST api/full-reset