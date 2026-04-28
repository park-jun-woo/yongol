//ff:func feature=cli-init type=util control=iteration dimension=1
//ff:what writeSkeletonFiles — materializes every skeleton file into the target directory

package cliinit

// writeSkeletonFiles materializes each file in the skeleton, rendering
// templates when needed.
func writeSkeletonFiles(targetDir string, data templateData) error {
	for _, f := range skeletonFiles() {
		if err := writeSkeletonFile(targetDir, data, f); err != nil {
			return err
		}
	}
	return nil
}
