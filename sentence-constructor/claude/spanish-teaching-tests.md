<test-cases>
    <case id="simple-1">
        <english>I eat bread.</english>
        <vocabulary>
            <word>
                <spanish>comer</spanish>
                <english>eat</english>
            </word>
            <word>
                <spanish>pan</spanish>
                <english>bread</english>
            </word>
        </vocabulary>
        <structure>[Subject] [Verb] [Object].</structure>
        <considerations>
            - Basic sentence with subject, verb, and object
            - Present tense form
            - Subject can be omitted in Spanish
        </considerations>
    </case>
    <case id="simple-2">
        <english>The book is red.</english>
        <vocabulary>
            <word>
                <spanish>libro</spanish>
                <english>book</english>
            </word>
            <word>
                <spanish>rojo</spanish>
                <english>red</english>
            </word>
        </vocabulary>
        <structure>[Subject] [Verb] [Adjective].</structure>
        <considerations>
            - Simple descriptor sentence
            - Uses ser vs estar
            - Adjective agreement required
        </considerations>
    </case>
</test-cases>

### 1.2 Compound Sentences
```xml
<test-cases>
    <case id="compound-1">
        <english>I eat bread and drink water.</english>
        <vocabulary>
            <word>
                <spanish>comer</spanish>
                <english>eat</english>
            </word>
            <word>
                <spanish>pan</spanish>
                <english>bread</english>
            </word>
            <word>
                <spanish>beber</spanish>
                <english>drink</english>
            </word>
            <word>
                <spanish>agua</spanish>
                <english>water</english>
            </word>
        </vocabulary>
        <structure>[Subject] [Verb1] [Object1] y [Verb2] [Object2].</structure>
        <considerations>
            - Compound sentence with two actions
            - Subject shared between clauses
            - Uses "y" for connection
        </considerations>
    </case>
</test-cases>

### 1.3 Complex Sentences
```xml
<test-cases>
    <case id="complex-1">
        <english>Because it's hot, I drink water.</english>
        <vocabulary>
            <word>
                <spanish>calor</spanish>
                <english>hot</english>
            </word>
            <word>
                <spanish>beber</spanish>
                <english>drink</english>
            </word>
            <word>
                <spanish>agua</spanish>
                <english>water</english>
            </word>
        </vocabulary>
        <structure>[Porque clause] [Subject] [Verb] [Object].</structure>
        <considerations>
            - Cause and effect relationship
            - Uses "porque" for "because"
            - Weather expression with "hacer"
        </considerations>
    </case>
</test-cases>

## 2. Vocabulary Edge Cases
### 2.1 Multiple Meanings
```xml
<vocabulary-test>
    <word>
        <spanish>estar</spanish>
        <meanings>
            <meaning>to be (temporary state)</meaning>
            <meaning>to be (location)</meaning>
            <meaning>to be (condition)</meaning>
        </meanings>
        <test-sentences>
            <sentence>How are you feeling?</sentence>
            <sentence>Where is the book?</sentence>
            <sentence>The food is ready.</sentence>
        </test-sentences>
    </word>
</vocabulary-test>

### 2.2 Regular/Irregular Verb Pairs
```xml
<vocabulary-test>
    <pair>
        <regular>
            <spanish>hablar</spanish>
            <english>to speak</english>
        </regular>
        <irregular>
            <spanish>decir</spanish>
            <english>to say/tell</english>
        </irregular>
        <test-sentences>
            <sentence>I speak Spanish.</sentence>
            <sentence>I tell you the truth.</sentence>
        </test-sentences>
    </pair>
</vocabulary-test>

## 3. State Transition Tests
### 3.1 Valid Transitions
```xml
<transition-test>
    <scenario id="setup-to-attempt">
        <initial-state>Setup</initial-state>
        <input>Yo como pan.</input>
        <expected-state>Attempt</expected-state>
        <validation>
            - Input contains Spanish text
            - No question marks
            - Contains vocabulary from setup
        </validation>
    </scenario>
    <scenario id="attempt-to-clues">
        <initial-state>Attempt</initial-state>
        <input>How do I use prepositions?</input>
        <expected-state>Clues</expected-state>
        <validation>
            - Input is a question
            - References grammar concept
            - Related to previous attempt
        </validation>
    </scenario>
</transition-test>

## 4. Teaching Scenario Tests
### 4.1 Common Mistakes
```xml
<teaching-test>
    <scenario id="ser-estar-mistake">
        <student-attempt>La comida es caliente.</student-attempt>
        <error>Incorrect use of ser instead of estar for temporary state</error>
        <expected-guidance>
            - Acknowledge attempt
            - Explain ser vs estar without giving answer
            - Encourage new attempt
        </expected-guidance>
    </scenario>
    <scenario id="conjugation-mistake">
        <student-attempt>Yo sabo la respuesta.</student-attempt>
        <error>Incorrect irregular verb conjugation</error>
        <expected-guidance>
            - Point out verb irregularity
            - Review present tense irregular verbs
            - Encourage correction
        </expected-guidance>
    </scenario>
</teaching-test>

## 5. Validation Criteria
### 5.1 Response Scoring
```xml
<scoring-criteria>
    <category name="vocabulary-table">
        <criteria>
            - Contains all necessary words (2 points)
            - Correct formatting (2 points)
            - Infinitive forms only (2 points)
            - No preposition inclusion (2 points)
            - Appropriate difficulty level (2 points)
        </criteria>
    </category>
    <category name="sentence-structure">
        <criteria>
            - Clear bracketed format (2 points)
            - No conjugations shown (2 points)
            - Appropriate for level (2 points)
            - Matches example patterns (2 points)
            - No prepositions included (2 points)
        </criteria>
    </category>
</scoring-criteria>

## 6. Documentation Improvements
[Previous version control and cross-references sections remain the same but with Spanish-specific references]