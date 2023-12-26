module default {
    type User {
        required property name -> str;
        required property email -> str {
            constraint exclusive;
        }
        required property password -> str;
        multi link following -> User;
        multi link followers -> User;
    }

    type Post {
        required property title -> str;
        required property content -> str;
        required link author -> User;
    }

    type Notification {
        required property type -> NotificationType;
        required property message -> str;
        required link user -> User;
    }

    scalar type NotificationType extending enum<FOLLOWER>;
}