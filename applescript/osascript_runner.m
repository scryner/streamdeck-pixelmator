#import "osascript_runner.h"

int run_osascript(const char *script, const char** result_message, const char **error_message) {
    NSString *scriptString = [NSString stringWithUTF8String:script];
    NSAppleScript *appleScript = [[NSAppleScript alloc] initWithSource:scriptString];
    NSDictionary *error;
    NSAppleEventDescriptor *result = [appleScript executeAndReturnError:&error];

    if (result) {
//         NSLog(@"Script executed successfully: %@", result);
        NSString *resultString = [result stringValue];
        *result_message = [resultString UTF8String];
        return 0;
    } else {
//         NSLog(@"Error executing script: %@", error);
        NSString *errorString = [error objectForKey:NSAppleScriptErrorMessage];
        *error_message = [errorString UTF8String];
        return -1;
    }
}
